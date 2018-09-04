package main

import "../myPackage"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	myPackage.PrintSlice("Slice", s)

	s = s[:0]
	myPackage.PrintSlice("Slice", s)

	s = s[2:4]
	myPackage.PrintSlice("Slice", s)

	s = s[:4]
	myPackage.PrintSlice("Slice", s)
}
