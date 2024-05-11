package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
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

func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
