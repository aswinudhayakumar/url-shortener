package helper

import (
	crand "crypto/rand"
	"crypto/sha256"
	"math/rand"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateRandomStringWithKey(key string, length int) (string, error) {
	b := make([]byte, length)
	if _, err := crand.Read(b); err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(key))

	r := rand.New(rand.NewSource(int64(hash[0])))
	for i := 0; i < length; i++ {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b), nil
}
