package main

import (
	"../myPackage"
	"fmt"
)

func main() {
  	i, j := 42, myPackage.GetIntFromInput("Input number: ", 5000)

  	p := &i
  	fmt.Println(*p)
  	*p = 21
  	fmt.Println(i)

  	p = &j
  	*p = *p / 37
  	fmt.Println(j)
}
