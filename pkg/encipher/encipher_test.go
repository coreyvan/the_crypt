package encipher

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncipherDecipher(t *testing.T) {
	plaintext := "Hello, World!"

	key := make([]byte, 32)
	_, err := rand.Read(key)
	require.NoError(t, err)

	ciphertext, err := Encrypt(key, []byte(plaintext))
	require.NoError(t, err)

	decrypted, err := Decrypt(key, []byte(ciphertext))
	require.NoError(t, err)

	assert.Equal(t, plaintext, decrypted)
}
