package main

import (
	"../myPackage"
	"fmt"
)

var i, j = 66, myPackage.GetIntFromInput("Please input number: ", 100)

func main() {
  	var c, python, java = i > j, j > i, "no!"
  	fmt.Println(i, j, c, python, java)
}
