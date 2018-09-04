package main

import (
	"../myPackage"
	"fmt"
	"math"
)

func PythagorasSquareRoot(x float64) float64 {
	squareRoot := float64(1)

	for i := 1; i < int(x)+3; i++ {
		squareRoot = 0.5 * (squareRoot + x/squareRoot)
	}

	return squareRoot
}

func main() {
	number := float64(myPackage.GetIntFromInput("Input number: ", 100))

	fmt.Printf("Pythagoras square root of %v is %v\n", number, PythagorasSquareRoot(number))
	fmt.Printf("Base func square root of %v is %v", number, math.Sqrt(number))
}
