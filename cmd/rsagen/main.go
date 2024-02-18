package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"os"
)

func main() {
	privateKey, err := generateKey(2048)
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	derBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	var out bytes.Buffer
	if err := pem.Encode(io.Writer(&out), &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derBytes,
	}); err != nil {
		log.Fatalf("Failed to write private key: %v", err)
	}

	if err := os.WriteFile("private.pem", out.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("Failed to write private key to file: %v", err)
	}

	out.Reset()
	derBytes, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Failed to marshal public key: %v", err)
	}

	if err := pem.Encode(io.Writer(&out), &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}); err != nil {
		log.Fatalf("Failed to write public key to file: %v", err)
	}

	if err := os.WriteFile("public.pem", out.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("Failed to write public key to file: %v", err)
	}
}

func generateKey(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
