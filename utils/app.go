package utils

import (
	"fmt"
	"math/rand"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateForgetToken(id int) string {

	suffix := fmt.Sprintf("%03d", id)
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b) + suffix

}
