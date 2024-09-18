package encipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(key []byte, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext = Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	enc := cipher.NewCBCEncrypter(block, iv)

	enc.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext), nil
}

func Decrypt(key []byte, ciphertext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	dec := cipher.NewCBCDecrypter(block, iv)

	dec.CryptBlocks(plaintext, ciphertext)
	plaintext = Unpad(plaintext, aes.BlockSize)

	return string(plaintext), nil

}

func Pad(data []byte, blockSize int) []byte {
	n := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(n)}, n)
	return append(data, padding...)
}

func Unpad(data []byte, blockSize int) []byte {
	n := int(data[len(data)-1])
	return data[:len(data)-n]
}
