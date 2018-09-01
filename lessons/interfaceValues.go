package main

import (
  	"fmt"
  	"math"
)

type I2 interface {
  	M()
}

type T2 struct {
  	S string
}


func (t *T2) M() {
  	fmt.Println(t.S)
}


type F2 float64


func (f F2) M() {
  	fmt.Println(f)
}


func main() {
  	var i I2

  	i = &T2{"Hello"}
  	describe(i)
  	i.M()

  	i = F2(math.Pi)
  	describe(i)
  	i.M()
}


func describe(i I2) {
  	fmt.Printf("(%v, %T)\n", i, i)
}
