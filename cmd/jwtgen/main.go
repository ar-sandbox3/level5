package main

import (
	"flag"
	"log"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var (
	key       string
	symmetric bool
)

func main() {
	flag.StringVar(&key, "key", "private.pem", "Path to private key file or simmetryc key")
	flag.BoolVar(&symmetric, "symmetric", false, "Use symmetric key")

	flag.Parse()

	var jwkKey jwk.Key
	var err error

	if symmetric {
		jwkKey, err = jwk.FromRaw([]byte(key))
		if err != nil {
			log.Fatalf("Failed to parse symmetric key: %v", err)
		}

	} else {
		keyBytes, err := os.ReadFile("private.pem")
		if err != nil {
			log.Fatalf("Failed to read private key file: %v", err)
		}

		jwkKey, err = jwk.ParseKey(keyBytes, jwk.WithPEM(true))
		if err != nil {
			log.Fatalf("Failed to parse private key: %v", err)
		}
	}

	token, err := generateToken(jwkKey, "https://sandbox.hsi.id", symmetric)
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	if err := os.WriteFile("token.jwt", token, os.ModePerm); err != nil {
		log.Fatalf("Failed to write token to file: %v", err)
	}
}

func generateToken(key jwk.Key, issuer string, symmetric bool) ([]byte, error) {
	token, err := jwt.NewBuilder().
		Subject("dio@hsi.id").
		Claim("scope", "read:foo").
		Issuer(issuer).
		Build()

	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if symmetric {
		return jwt.Sign(token, jwt.WithKey(jwa.HS256, key))
	}

	return jwt.Sign(token, jwt.WithKey(jwa.RS256, key))
}
