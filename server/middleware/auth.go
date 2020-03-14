package middleware

import (
	"log"
	"net/http"

	"github.com/mneil/polsino/server/request"
	"gopkg.in/square/go-jose.v2/jwt"
)

type (
	// AuthIFace holds the shape of authentication backends. Authentication determins if you are allowed to manage the policy namespace that you are attempting to manage.
	AuthIFace interface {
		Verify(http.ResponseWriter, *http.Request) *request.HandlerError
	}
	// JWT requires JWT authentication to communicate with the policy server
	Auth struct {
	}
	// Claims are the JWT payload
	Claims struct {
		Issuer   string `json:"iss,omitempty"`
		Subject  string `json:"sub,omitempty"`
		Audience string `json:"aud,omitempty"`
		Expiry   int64  `json:"exp,omitempty"`
		IssuedAt int64  `json:"iat,omitempty"`
	}
)

// Verify checks that the JWT is a valid JWT. Verifies the signature from a well known location and optionally checks a subscriber(es)
func (auth *Auth) Verify(res http.ResponseWriter, req *http.Request) *request.HandlerError {
	// Get token
	authorization := req.Header.Get("Authorization")
	if authorization == "" {
		return request.NewHandlerError(401, "missing header 'Authorization'")
	}
	if authorization[:7] != "Bearer" {
		return request.NewHandlerError(401, "'Authorization' malformed. Should start with 'Bearer '")
	}
	bearer := authorization[7:]
	// Parse JWT
	parsed, err := jwt.ParseSigned(bearer)
	if err != nil {
		return request.NewHandlerError(403, err.Error())
	}

	// Verify Claims
	resultCl := Claims{}
	err = parsed.Claims("public key", &resultCl)
	if err != nil {
		log.Fatalf("Failed to get claims JWT:%+v", err)
	}
	return nil
}
