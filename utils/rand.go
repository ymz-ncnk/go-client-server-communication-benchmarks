package utils

import "math/rand"

const MaxStringLength = 1007

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomBool(r *rand.Rand) bool {
	return r.Intn(2) == 1
}

func RandomInt64(r *rand.Rand) int64 {
	return r.Int63()
}

func RandomInt(r *rand.Rand) int {
	return r.Int()
}

func RandomString(r *rand.Rand) string {
	var (
		length = r.Intn(MaxStringLength)
		b      = make([]rune, length)
	)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func RandomFloat64(r *rand.Rand) float64 {
	return r.NormFloat64()
}
