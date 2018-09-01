package main

import (
  	"fmt"
  	"math/cmplx"
)

var (
  	ToBe bool
  	MaxInt uint64 = 1<<64 - 1
  	z = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T\t\t\tValue: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T\t\tValue: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T\tValue: %v\n", z, z)
}
