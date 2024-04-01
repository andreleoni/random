package random

import (
	"math/rand"
)

var numbersRunes = []rune("1234567890")

// Numbers returns a numbers.
func Numbers(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(numbersRunes))]
	}

	return string(b)
}
