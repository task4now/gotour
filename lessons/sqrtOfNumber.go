package main

import (
	"../myPackage"
    "fmt"
	"math"
)

func main() {
    number := myPackage.GetIntFromInput("Please, input number: ", 100)
    fmt.Printf("Now you have %g problems", math.Sqrt(float64(number)))
}
