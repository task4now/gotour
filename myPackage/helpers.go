package myPackage

import "fmt"


func PrintSlice(str string, slice []int) {
	fmt.Printf("%s: len=%d; cap=%d; %v\n", str, len(slice), cap(slice), slice)
}


func DescribeInterface(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}
