package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Ridj Forever", 17}
	z := Person{"Sergey Romanov", 59}
	fmt.Println(a, z)
}
