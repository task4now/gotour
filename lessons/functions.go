package main

import (
	"../myPackage"
  	"fmt"
	"math"
)


func compute(fn func(float64, float64) float64) float64 {
  	return fn(4, 3)
}


func main() {
  	hypot := func(x, y float64) float64 {
    	return math.Sqrt(x*x + y*y)
  	}

  	catet1, catet2 := 	float64(myPackage.GetIntFromInput("First catet: ", 10)),
  						float64(myPackage.GetIntFromInput("Second catet", 10))

  	fmt.Println(hypot(catet2, catet1))


  	fmt.Println(compute(hypot))
  	fmt.Println(compute(math.Pow))
}
