package util

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyz"

func GetRandomChars(number int) string {
	b := make([]byte, number)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)

}
