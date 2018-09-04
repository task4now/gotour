package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := localFloat1(-math.Sqrt2)
	v := localVertex6{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}

type localFloat1 float64

func (f localFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type localVertex6 struct {
	X, Y float64
}

func (v *localVertex6) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
