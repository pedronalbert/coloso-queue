package queue

import (
  "errors"
)
// ErrNotMoreEntries
var ErrNotMoreEntries = errors.New("There is not more entries")

// ErrNotAdded
var ErrNotAdded = errors.New("Can't added to queue")

// ErrDuplicateEntry
var ErrDuplicateEntry = errors.New("Entry is already in queue")
