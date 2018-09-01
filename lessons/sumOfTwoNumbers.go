package main

import (
	"../myPackage"
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	numberOne := myPackage.GetIntFromInput("Please input first number: ", 1000)
	numberTwo := myPackage.GetIntFromInput("Please input second number: ", 1000)

	sum := add(numberOne, numberTwo)
	fmt.Printf("\n\t%d + %d = %d\n", numberOne, numberTwo, sum)
}
