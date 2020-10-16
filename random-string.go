package go_str_random

import (
	"math/rand"
)

// Alphabet
const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomString(n int) string {

	b := make([]byte, n)

	for i := range b {
		b[i] = Alphabet[rand.Intn(len(Alphabet))]
	}

	return string(b)
}
