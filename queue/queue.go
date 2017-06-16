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

// JSONEntry - estructura de una entrada en la cola
type JSONEntry struct {
  FetchType string `json:"fetchType"`
  FetchID string `json:"fetchId"`
}

// GetLength - devuelve la longitud actual de la cola
func GetLength() int64 {
  length, _ := redis.Client.LLen(queueName).Result()

  return length
}

// GetNextEntry - devuelve la proxima entrada en la cola
func GetNextEntry() (JSONEntry, error) {
  var jsonEntry JSONEntry

  textEntry, err := redis.Client.LPop(queueName).Result()


  if err != nil {
    log.Info("Not more entries in queue")
    return jsonEntry, ErrNotMoreEntries
  }

  log.Debugf("Entry loaded from redis: %s", textEntry)

  err = json.Unmarshal([]byte(textEntry), &jsonEntry)

  if err != nil {
    log.Error("Can't decode the JSON")
    return jsonEntry, err
  }

  return jsonEntry, nil
}

// GetAllEntries - Devuelve todas las entradas de la cola
func GetAllEntries() ([]JSONEntry, error) {
  entries := []JSONEntry{}
  entriesString, err := redis.Client.LRange(queueName, 0, -1).Result()

  if err != nil {
    log.Error("Can't load entries list", err)

    return entries, err
  }



  for _, entryString := range entriesString {
    var entry JSONEntry

    err = json.Unmarshal([]byte(entryString), &entry)

    if (err == nil) {
      entries = append(entries, entry)
    } else {
      log.Error("Can't encode JSON: ", entryString)
    }
  }

  return entries, nil
}

// AddEntry - Agregar entrada al final de la cola
func AddEntry(entry JSONEntry) (queuePosition int64, err error) {
  entryBytes, err := json.Marshal(entry)
  entryString := string(entryBytes)

  if err != nil {
    log.Error("Can't create JSON entry")
    return -1, err
  }

  result := redis.Client.LPush(queueName, entryString)

  if result.Err() != nil {
    log.Error("Can't add entry to queue")

    return -1, result.Err()
  }

  queuePos, _ := result.Result()
  log.Info("Entry added to position: ", queuePos)

  return queuePos, nil
}
