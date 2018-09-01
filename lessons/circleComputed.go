package main

import (
	"../myPackage"
	"fmt"
	"math"
)

func main() {

	radius := float64(myPackage.GetIntFromInput("Please, input radius of circle: ", 50))

	Pi := math.Pi
	perimeter := 2 * Pi * radius
	square := Pi * radius * radius

	fmt.Printf("Perimeter of circle is %.2f and square is %.2f", perimeter, square)
}
