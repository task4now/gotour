package main

import (
	"../myPackage"
	"fmt"
)

func main() {
	v := myPackage.Vertex{X: 1, Y: 2}
	v.X = myPackage.GetIntFromInput("Input number: ", 100)
	fmt.Println("New Vertex is", v)
}
