package utils

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString() (s string) {
	b := make([]byte, rand.Intn(MaxStringLength))
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
