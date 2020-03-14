package request

import (
	"fmt"
	"net/http"
)

type (
	// HandlerError is an error with a response code
	HandlerError struct {
		Err  string
		Code int
	}
	// Handler is an http handler func that returns an error for chain processing
	Handler func(http.ResponseWriter, *http.Request) *HandlerError
)

func (e *HandlerError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Err)
}

// NewHandlerError creates a new HandlerError
func NewHandlerError(code int, err string) *HandlerError {
	return &HandlerError{
		Err:  err,
		Code: code,
	}
}

// Handlers processes n HandlerFunc until a response is returned
func Handlers(handlers []Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			if err := handler(w, r); err != nil {
				http.Error(w, err.Error(), err.Code)
				break
			}
		}
	}
}
