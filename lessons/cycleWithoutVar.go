package main

import (
	"../myPackage"
	"fmt"
	"math"
)

func main() {
  	multi := float64(myPackage.GetIntFromInput("Input final number: ", 100))

  	for ; multi <= math.Pow10(50); {
    	multi *= multi
  	}

  	fmt.Println("So, we came to", multi)
}
