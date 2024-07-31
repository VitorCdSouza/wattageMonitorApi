package util

import (
	"math/rand"
)

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

