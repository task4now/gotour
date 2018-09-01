package main

import (
	"../myPackage"
	"fmt"
)

func main() {
  	sum := 0

  	count := myPackage.GetIntFromInput("Input final number: ", 100)

  	for i := 0; i <= count; i++ {
    	sum += i
  	}

  	fmt.Printf("Sum of numbers from zero to %d is %d.", count, sum)
}
