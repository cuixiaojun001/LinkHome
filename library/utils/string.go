package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	result := make([]byte, 28)
	for i := range result {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[val.Int64()]
	}
	return string(result)
}
