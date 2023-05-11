package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encrypt(t *testing.T) {
	t.Run("encrypts and decrypts value", func(t *testing.T) {
		key := "passphrase"
		value := "encrypt this"
		encrypted := encrypt(key, value)
		decrypted := decrypt(key, encrypted)
		assert.Equal(t, value, decrypted)
	})

}
