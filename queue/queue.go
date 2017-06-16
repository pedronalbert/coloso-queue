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
    log.Info("No hay mas entradas en la cola")
    return jsonEntry, ErrNotMoreEntries
  }

  log.Debugf("Obtenida entrada de redis: %s", textEntry)

  err = json.Unmarshal([]byte(textEntry), &jsonEntry)

  if err != nil {
    log.Critical("No se ha podido decodificar el JSON: %s", err)
    return jsonEntry, err
  }

  return jsonEntry, nil
}

// GetAllEntries - Devuelve todas las entradas de la cola
func GetAllEntries() ([]JSONEntry, error) {
  
}

// AddEntry - Agregar entrada al final de la cola
func AddEntry(entry JSONEntry) (queuePosition int64, err error) {
  entryBytes, err := json.Marshal(entry)
  entryString := string(entryBytes)

  if err != nil {
    log.Error("No se ha podido crear el JSON")
    return -1, err
  }

  result := redis.Client.LPush(queueName, entryString)

  if result.Err() != nil {
    log.Error("No se ha podido añadir a la cola")

    return -1, result.Err()
  }

  queuePos, _ := result.Result()
  log.Info("Entrada añádida a la cola en la posicion ", queuePos)

  return queuePos, nil
}
