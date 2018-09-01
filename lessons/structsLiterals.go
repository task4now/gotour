package main

import (
	"../myPackage"
	"fmt"
)

var (
  	v1 = myPackage.Vertex{X: 1, Y: 2}
  	v2 = myPackage.Vertex{X: 1}
  	v3 = myPackage.Vertex{}
  	p  = &myPackage.Vertex{X: 1, Y: 2}
)

func main() {
  	fmt.Println(v1, p, v2, v3)
}
