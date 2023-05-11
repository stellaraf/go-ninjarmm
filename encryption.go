package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func encrypt(passphrase, plaintext string) string {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

func decrypt(passphrase, ciphertext string) string {
	arr := strings.Split(ciphertext, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}
