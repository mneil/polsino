package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"testing"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func generateJWT() (string, string) {
	// For testing create the RSA key pair in the code
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	} // create Square.jose signing key
	key := jose.SigningKey{Algorithm: jose.RS256, Key: privKey}
	// create a Square.jose RSA signer, used to sign the JWT
	var signerOpts = jose.SignerOptions{}
	signerOpts.WithType("JWT")
	rsaSigner, err := jose.NewSigner(key, &signerOpts)
	if err != nil {
		log.Fatalf("failed to create signer:%+v", err)
	}
	// create an instance of Builder that uses the rsa signer
	builder := jwt.Signed(rsaSigner)

	// public claims
	pubClaims := Claims{
		Issuer:   "issuer1",
		Subject:  "subject1",
		ID:       "id1",
		Audience: jwt.Audience{"aud1", "aud2"},
		IssuedAt: time.Now().Unix(),
		Expiry:   time.Now().Unix() + 1000,
	}

	builder = builder.Claims(pubClaims)

	// validate all ok, sign with the RSA key, and return a compact JWT
	jwt, err := builder.CompactSerialize()
	if err != nil {
		log.Fatalf("failed to create JWT:%+v", err)
	}
	return jwt, ""
}

func TestVerify(t *testing.T) {
	// Verify()
}
