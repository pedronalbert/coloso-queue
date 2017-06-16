package riot

import (
  "errors"
)

// ErrBadRequest = Error 400 en requests
var ErrBadRequest = errors.New("Error 400 - Bad Request")

// ErrUnauthorized
var ErrUnauthorized = errors.New("Error 401 - Unauthorized")

// ErrForbidden
var ErrForbidden = errors.New("Error 403 - Forbidden")

// ErrNotFound - Error 404 en requests
var ErrNotFound = errors.New("Error 404 - Not Found")

// ErrServerError - Errores 5XX en requests
var ErrServerError = errors.New("Error 5XX - Server Error")
