package main

import (
	"fmt"
	"math"
)

type localVortex4 struct {
	X, Y float64
}

func (v localVortex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v localVortex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := localVortex4{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &localVortex4{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
