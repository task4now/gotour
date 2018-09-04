package main

import (
	"../myPackage"
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	x := float64(myPackage.GetIntFromInput("Input number: ", 1000000))
	y := x - math.Pow10(6)

	fmt.Printf("Square root of %v is %v and sqrt of %v is %v.", x, sqrt(x), y, sqrt(y))
}
