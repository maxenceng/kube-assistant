package utils

import (
	"math/rand"
	"time"
)

func RandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, 10)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}
