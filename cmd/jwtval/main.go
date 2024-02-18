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
		keyBytes, err := os.ReadFile(key)
		if err != nil {
			log.Fatalf("Failed to read private key file: %v", err)
		}

		jwkKey, err = jwk.ParseKey(keyBytes, jwk.WithPEM(true))
		if err != nil {
			log.Fatalf("Failed to parse private key: %v", err)
		}
	}

	jwtBytes, err := os.ReadFile("token.jwt")
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	algorithm := jwa.RS256
	if symmetric {
		algorithm = jwa.HS256
	}
	token, err := jwt.Parse(jwtBytes, jwt.WithKey(algorithm, jwkKey))
	if err != nil {
		log.Fatalf("Failed to parse JWT: %v", err)
	}

	if err := jwt.Validate(token,
		jwt.WithIssuer("https://sandbox.hsi.id"),
		jwt.WithClaimValue("scope", "read:foo"),
	); err != nil {
		log.Fatalf("Failed to validate JWT: %v", err)
	}
}
