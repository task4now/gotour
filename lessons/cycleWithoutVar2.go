package main

import (
	"../myPackage"
	"fmt"
)

func main() {
  	sum := myPackage.GetIntFromInput("Input number: ", 100)

  	for sum < 10000 {
    	sum += sum
  	}

  	fmt.Println(sum)
}
