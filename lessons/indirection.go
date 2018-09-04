package main

import "fmt"

type localVertex3 struct {
	X, Y float64
}

func (v *localVertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *localVertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := localVertex3{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &localVertex3{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
