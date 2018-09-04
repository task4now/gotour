package main

import (
	"../myPackage"
	"fmt"
	"math"
)

func Abs(v myPackage.VertexFloat) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := myPackage.VertexFloat{3, 4}
	fmt.Println(Abs(v))
}
