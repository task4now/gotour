package main

import (
	"../myPackage"
  	"fmt"
	"math"
)

func main() {
	var x, y =  myPackage.GetIntFromInput("Input length of first catet: ", 10),
				myPackage.GetIntFromInput("Input length of second catet: ", 10)

	var hypotenuse = uint(math.Sqrt(float64(x*x + y*y)))

  	fmt.Printf("Hypotenuse of triangle with catets %d and %d is %d", x, y, hypotenuse)
}
