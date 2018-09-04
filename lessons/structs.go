package main

import (
	"../myPackage"
	"fmt"
)

// For structure creation see ../myPackage/vertex.go

func main() {
	numberOne, numberTwo := myPackage.GetRandomInt(100), myPackage.GetRandomInt(100)
	fmt.Println(myPackage.Vertex{X: numberOne, Y: numberTwo})
}
