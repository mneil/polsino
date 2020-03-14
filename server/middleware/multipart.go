package middleware

import (
	"net/http"

	"github.com/mneil/polsino/server/request"
)

// Multipart handles parsing a multipart request
func Multipart(http.ResponseWriter, *http.Request) *request.HandlerError {
	return nil
}
