package authenticator

import (
	"github.com/ar-sandbox3/level5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// Authenticator is an authenticator.
type Authenticator struct {
	Issuer     string
	PrivateKey jwk.Key
	PublicKey  jwk.Key
}

// NewAuthenticator returns an initialized authenticator.
func NewAuthenticator(issuer string) (*Authenticator, error) {
	privateKeyJWK, err := jwk.ParseKey(level5.Private, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	publicKeyJWK, err := jwk.ParseKey(level5.Public, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	return &Authenticator{
		Issuer:     issuer,
		PrivateKey: privateKeyJWK,
		PublicKey:  publicKeyJWK,
	}, nil
}

// Validate validates a token.
func (a Authenticator) Validate(token []byte) error {
	jwtToken, err := jwt.Parse(token, jwt.WithKey(jwa.RS256, a.PublicKey))
	if err != nil {
		return err
	}

	return jwt.Validate(jwtToken, jwt.WithIssuer(a.Issuer))
}

// Generate generates a token.
func (a Authenticator) Generate(subject string) ([]byte, error) {
	token, err := jwt.NewBuilder().
		Subject(subject).
		Issuer(a.Issuer).
		Build()

	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return jwt.Sign(token, jwt.WithKey(jwa.RS256, a.PrivateKey))
}
