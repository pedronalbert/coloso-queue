package queue

import (
  "errors"
  "encoding/json"
  "coloso-queue/clients/redis"
  "github.com/op/go-logging"
)

var log = logging.MustGetLogger("queue")

// Name of the queue
const queueName = "coloso_queue"

// ErrNotMoreEntries
var ErrNotMoreEntries = errors.New("No hay mas entradas en cola")

// ErrNotAdded
var ErrNotAdded = errors.New("No se ha podido añadir a la cola")

// ErrDuplicateEntry
var ErrDuplicateEntry = errors.New("La entrada ya se encuentra en cola")

// JSONEntry - estructura de una entrada en la cola
type JSONEntry struct {
  FetchType string `json:"fetchType"`
  FetchID string `json:"fetchId"`
}

// GetLength - Devuelve la longitud actual de la cola
func GetLength() int64 {
  length, _ := redis.Client.LLen(queueName).Result()

  return length
}

// GetNextEntry - Devuelve la primera entrada en la cola
func GetNextEntry() (JSONEntry, error) {
  var jsonEntry JSONEntry

  textEntry, err := redis.Client.LPop(queueName).Result()


  if err != nil {
    log.Info("No hay mas entradas en cola")
    return jsonEntry, ErrNotMoreEntries
  }

  log.Debugf("Entrada leida desde redis: %s", textEntry)

  err = json.Unmarshal([]byte(textEntry), &jsonEntry)

  if err != nil {
    log.Error("No se ha podido decodificar decodifiar el JSON")
    return jsonEntry, err
  }

  return jsonEntry, nil
}

// GetAllEntries - Devuelve todas las entradas de la cola
func GetAllEntries() ([]JSONEntry, error) {
  entries := []JSONEntry{}
  entriesString, err := redis.Client.LRange(queueName, 0, -1).Result()

  if err != nil {
    log.Error("No se ha podido cargar la lista de entradas", err)

    return entries, err
  }

  for _, entryString := range entriesString {
    var entry JSONEntry

    err = json.Unmarshal([]byte(entryString), &entry)

    if (err == nil) {
      entries = append(entries, entry)
    } else {
      log.Error("No se ha podido codificar el JSON", entryString)
    }
  }

  return entries, nil
}

func hasDuplicateEntry(entries []JSONEntry, entryToFind JSONEntry) bool {
  for _, entry := range entries {
    if entry == entryToFind {
      return true
    }
  }

  return false
}

// Enqueue - Agregar entrada al final de la cola
func Enqueue(entry JSONEntry) (queuePosition int64, err error) {
  entryBytes, err := json.Marshal(entry)
  entryString := string(entryBytes)
  isEnqueued, err := IsEnqueued(entry)

  if isEnqueued || err != nil {
    log.Info("La entrada ya se encuentra en cola")
    return -1, ErrDuplicateEntry
  }

  if err != nil {
    log.Error("No se ha podido crear el JSON")
    return -1, err
  }

  result := redis.Client.RPush(queueName, entryString)

  if result.Err() != nil {
    log.Error("No se ha podido agregar la entrada a la cola")

    return -1, result.Err()
  }

  queuePos, _ := result.Result()
  log.Info("Entrada agregada a la cola en la posicion: ", queuePos)

  return queuePos, nil
}

// IsEnqueued - Chequea si la entrada que recibe ya se encuentra en cola
func IsEnqueued(entry JSONEntry) (bool, error) {
  entries, err := GetAllEntries()

  if err != nil {
    return true, err
  }

  return hasDuplicateEntry(entries, entry), nil
}
