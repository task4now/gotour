package main

import (
	"fmt"
	"math"
)

type localVertex2 struct {
	X, Y float64
}

func (v localVertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *localVertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := localVertex2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
