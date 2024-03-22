package web

import "errors"

// shutdownError is a type used to help with the graceful termination of the service.
type shutdownError struct {
	Message string
}

func NewShutdownError(message string) error {
	return &shutdownError{Message: message}
}

func (se *shutdownError) Error() string {
	return se.Message
}

func IsShutdown(err error) bool {
	var se *shutdownError
	return errors.As(err, &se)
}
