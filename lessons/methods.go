package main

import (
	"fmt"
	"math"
)

type localVertex1 struct {
	X, Y float64
}


func (v localVertex1) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


func main() {
  	v := localVertex1{3, 4}
  	fmt.Println(v.Abs())
}
