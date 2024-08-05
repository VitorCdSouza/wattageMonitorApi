package util

import (
	"math"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const symbols = "!@#$%&"
const (
	domain1 = "@gmail.com"
	domain2 = "@hotmail.com"
	domain3 = "@outlook.com"
)

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail(n int) string {
	domains := []string{domain1, domain2, domain3}
	k := len(domains)
	first := RandomString(n)
	return first + domains[rand.Intn(k)]
}

func RandomPassword(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	l := len(symbols)

	half := math.Floor(float64(n / 2))
	for i := 0; i < (n - int(half)); i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	for i := 0; i < (n - int(half)); i++ {
		c := symbols[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}