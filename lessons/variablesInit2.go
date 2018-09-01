package main

import (
	"../myPackage"
	"fmt"
)

func main() {
  	var i, j = 75, 50
  	k := myPackage.GetIntFromInput("Please input number: ", 100)

  	c, python, java := i > k, j > k, "no!"

  fmt.Println(i, j, k, c, python, java)
}
