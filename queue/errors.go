package queue

import (
  "errors"
)
// ErrNotMoreEntries
var ErrNotMoreEntries = errors.New("No hay mas entradas en cola")

// ErrNotAdded
var ErrNotAdded = errors.New("No se ha podido añadir a la cola")

// ErrDuplicateEntry
var ErrDuplicateEntry = errors.New("La entrada ya se encuentra en cola")
