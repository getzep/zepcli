package jwttools

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	// Test case when secret is not provided
	_, err := NewJWT([]byte(""))
	assert.ErrorContains(t, err, "auth secret not provided")

	// Test case when a valid secret is provided
	secret := []byte("mysecret")
	token, err := NewJWT(secret)
	assert.NoError(t, err)

	// Check if token is not empty
	assert.True(t, strings.HasPrefix(token, "ey"))
}

func TestNewSecret(t *testing.T) {
	secret, err := NewSecret(10)
	assert.NoError(t, err)
	assert.NotEmpty(t, secret)
}
