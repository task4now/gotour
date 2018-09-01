package main

import (
	"../myPackage"
	"fmt"
)

func main() {
  	v := myPackage.Vertex{X: 1, Y: 2}
  	p := &v
  	p.X = 1e9

  	fmt.Println(v)
}
