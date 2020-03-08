package http

import (
	"net/http"
)

type (
	// HandlerError is an error with a response code
	HandlerError struct {
		error
		Code int
	}
	// Handler is an http handler func that returns an error for chain processing
	Handler func(http.ResponseWriter, *http.Request) *HandlerError
)

// RequestHandler processes n HandlerFunc until a response is returned
func RequestHandler(handlers []Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			if err := handler(w, r); err != nil {
				http.Error(w, err.Error(), err.Code)
				break
			}
		}
	}
}
