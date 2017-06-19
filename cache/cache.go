package cache

import "errors"
import "github.com/op/go-logging"

var log = logging.MustGetLogger("queue")

// ErrNotFound - Entity not found on DB
var ErrNotFound = errors.New("Entity not found on DB")
