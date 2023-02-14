package utils

import (
	"math/rand"
	"time"
)

// RandNum 随机数1-10
func RandNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}
