package utils

import (
	"crypto/rand"
	"encoding/base32"
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
