package main

import (
	"../myPackage"
)

func main() {
	var s []int
	myPackage.PrintSlice("Slice initial", s)

	s = append(s, 1)
	myPackage.PrintSlice("Slice single append", s)

	s = append(s, 2, 3, 4)
	myPackage.PrintSlice("Slice multiple append", s)
}
