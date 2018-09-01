package main

import (
	"../myPackage"
  	"fmt"
	"math"
)


func AbsFloat(v myPackage.VertexFloat) float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func Scale(v myPackage.VertexFloat, f float64) myPackage.VertexFloat {
  	v.X = v.X * f
  	v.Y = v.Y * f
  	return v
}


func main() {
  	v := myPackage.VertexFloat{3, 4}
  	v = Scale(v, 10)
  	fmt.Println(AbsFloat(v))
}
