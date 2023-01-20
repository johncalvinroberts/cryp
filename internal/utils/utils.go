package utils

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"fmt"
)

func RandomSecret(length int) string {
	var result string
	secret := make([]byte, length)
	gen, err := rand.Read(secret)
	if err != nil || gen != length {
		// error reading random, return empty string
		return result
	}
	var encoder = base32.StdEncoding.WithPadding(base32.NoPadding)
	result = encoder.EncodeToString(secret)
	return result
}

func EncryptMessage(secret string, message string) string {
	c, err := aes.NewCipher([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(message))
	c.Encrypt(msgByte, []byte(message))
	return hex.EncodeToString(msgByte)
}

func DecryptMessage(secret string, message string) string {
	txt, _ := hex.DecodeString(message)
	c, err := aes.NewCipher([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(txt))
	c.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg
}
