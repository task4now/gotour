package main

import (
	"../myPackage"
	"fmt"
)

func main() {
  	fmt.Println("Counting...")

  	final := myPackage.GetIntFromInput("Input number: ", 25)

  	for i := 0; i < final; i++ {
    	if i == 0 {
      		defer fmt.Print(i, ".")
    	} else {
      		defer fmt.Print(i, ", ")
    	}
  	}

  	fmt.Print("done: ")
}
