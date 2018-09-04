package myPackage

import (
	"math/rand"
	"time"
)

func GetRandomInt(count int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	return random.Intn(count)
}
