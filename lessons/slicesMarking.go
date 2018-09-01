package main

import "../myPackage"

func main() {
  a := make([]int, 5)

	myPackage.PrintSlice("Slice a", a)

  b := make([]int, 0, 5)
	myPackage.PrintSlice("Slice b", b)

  c := b[:2]
	myPackage.PrintSlice("Slice c", c)

  d := c[2:5]
	myPackage.PrintSlice("Slice d", d)
}
