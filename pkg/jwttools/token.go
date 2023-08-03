package jwttools

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

const SecretLength = 64

// NewJWT generates a JWT token using the given config.
func NewJWT(secret []byte) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("auth secret not provided")
	}

	tokenAuth := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := tokenAuth.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error generating auth token: %w", err)
	}

	return tokenString, nil
}

func NewSecret(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateJWT() error {
	fmt.Println("Generating a secret key that we'll use to sign your JWT tokens.\n\n" +
		"This should be set as the ZEP_AUTH_SECRET environment variable.\n\n" +
		"Do not use this key as your JWT token! It should be kept safe.\n" +
		"Anybody with access to this key can gain access to your Zep server.",
	)
	secret, err := NewSecret(SecretLength)
	if err != nil {
		return fmt.Errorf("unable to generate secret: %w", err)
	}
	fmt.Printf("\n\nSecret: %s\n\n", secret)
	fmt.Println("\nPress Enter once you have copied the secret to a safe place.")
	fmt.Scanln()

	fmt.Println("Generating a JWT token for use in your API calls.")

	key, err := NewJWT([]byte(secret))
	if err != nil {
		return fmt.Errorf("unable to generate JWT: %w", err)
	}

	fmt.Printf("\n\nJWT token: %s\n\n", key)

	fmt.Println("\nPress Enter once you have copied the JWT token to a safe place.")
	fmt.Scanln()

	fmt.Println("Complete. Reminder: Use the JWT Token as your API key and not the secret.")

	return nil
}
