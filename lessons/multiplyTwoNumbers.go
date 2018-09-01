package main

import (
	"../myPackage"
	"fmt"
)

func multiply(x, y int) int {
  	return x * y
}

func main() {
	numberOne := myPackage.GetIntFromInput("Please input first number: ", 1000)
	numberTwo := myPackage.GetIntFromInput("Please input second number: ", 1000)

	multi:= multiply(numberOne, numberTwo)
  	fmt.Printf("\n\t%d * %d = %d\n", numberOne, numberTwo, multi)
}
