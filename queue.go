package main

import (
  "strings"
  "encoding/json"
  "coloso-queue/clients/redis"
  "github.com/op/go-logging"
)

var log = logging.MustGetLogger("queue")

// QueueEntry - estructura de una entrada en la cola
type QueueEntry struct {
  FetchType string `json:"fetchType"`
  FetchID string `json:"fetchId"`
}

// Queue - stuct
type Queue struct {
  Name string
}

// NewQueue - create new queue
func NewQueue(region string) *Queue {
  region = strings.ToLower(region)

  return &Queue{
    Name: "coloso_queue_" + region,
  }
}

// GetLength - Devuelve la longitud actual de la cola
func (queue *Queue) GetLength() (int64, error) {
  length, err := redis.Client.LLen(queue.Name).Result()

  if err != nil {
    log.Error("No se ha podido obtener la longitud de la cola")
    log.Error(err)

    return 0, err
  }

  return length, nil
}

// GetNextEntry - Devuelve la primera entrada en la cola
func (queue *Queue) GetNextEntry() (QueueEntry, error) {
  var queueEntry QueueEntry

  textEntry, err := redis.Client.LPop(queue.Name).Result()

  if err != nil {
    log.Info("No hay mas entradas en cola")
    return queueEntry, ErrNotMoreEntries
  }

  log.Debugf("Entrada leida desde redis: %s", textEntry)

  err = json.Unmarshal([]byte(textEntry), &queueEntry)

  if err != nil {
    log.Error("No se ha podido decodificar decodifiar el JSON")
    return queueEntry, err
  }

  return queueEntry, nil
}

// GetAllEntries - Devuelve todas las entradas de la cola
func (queue *Queue) GetAllEntries() ([]QueueEntry, error) {
  entries := []QueueEntry{}
  entriesString, err := redis.Client.LRange(queue.Name, 0, -1).Result()

  if err != nil {
    log.Error("No se ha podido cargar la lista de entradas", err)

    return entries, err
  }

  for _, entryString := range entriesString {
    var entry QueueEntry

    err = json.Unmarshal([]byte(entryString), &entry)

    if (err == nil) {
      entries = append(entries, entry)
    } else {
      log.Error("No se ha podido codificar el JSON", entryString)
    }
  }

  return entries, nil
}

func hasDuplicateEntry(entries []QueueEntry, entryToFind QueueEntry) bool {
  for _, entry := range entries {
    if entry == entryToFind {
      return true
    }
  }

  return false
}

// Enqueue - Agregar entrada al final de la cola
func (queue *Queue) Enqueue(entry QueueEntry) (queuePosition int64, err error) {
  entryBytes, err := json.Marshal(entry)
  entryString := string(entryBytes)
  isEnqueued, err := queue.IsEnqueued(entry)

  if isEnqueued || err != nil {
    log.Info("La entrada ya se encuentra en cola")
    return -1, ErrDuplicateEntry
  }

  if err != nil {
    log.Error("No se ha podido crear el JSON")
    return -1, err
  }

  result := redis.Client.RPush(queue.Name, entryString)

  if result.Err() != nil {
    log.Error("No se ha podido agregar la entrada a la cola")

    return -1, result.Err()
  }

  queuePos, _ := result.Result()
  log.Info("Entrada agregada a la cola en la posicion: ", queuePos)

  return queuePos, nil
}

// IsEnqueued - Chequea si la entrada que recibe ya se encuentra en cola
func (queue *Queue) IsEnqueued(entry QueueEntry) (bool, error) {
  entries, err := queue.GetAllEntries()

  if err != nil {
    return true, err
  }

  return hasDuplicateEntry(entries, entry), nil
}
