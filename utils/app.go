package utils

import (
	"fmt"
	"math/rand"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateForgetToken(id uint) string {
	prefix := fmt.Sprintf("%03d", id)
	return generateRandomString(8) + prefix
}

func GenerateVerifToken(id uint) string {
	idStr := fmt.Sprintf("%03d", id)
	return generateRandomString(5) + idStr + generateRandomString(5)
}

func generateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
