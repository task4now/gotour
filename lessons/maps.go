package main

import (
	"../myPackage"
	"fmt"
)

var m1 map[string]myPackage.VertexMap

func main() {
	m1 = make(map[string]myPackage.VertexMap)

	m1["Bell Labs"] = myPackage.VertexMap{
		40.68433, -74.39967,
	}

	fmt.Println(m1["Bell Labs"])
}
