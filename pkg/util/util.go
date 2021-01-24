package util

import (
	crand "crypto/rand"
	"math/rand"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := crand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
