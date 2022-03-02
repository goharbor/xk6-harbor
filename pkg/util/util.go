package util

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"os"
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

func GetEnv(key string, defaults ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	panic(fmt.Errorf("%s environment is required", key))
}
