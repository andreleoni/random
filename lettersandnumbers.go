package random

import (
	"math/rand"
)

var letterAndNumberRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// LettersAndNumbers returns numbers and letters on strings.
func LettersAndNumbers(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterAndNumberRunes))]
	}

	return string(b)
}
