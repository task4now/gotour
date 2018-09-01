package main

import (
	"../myPackage"
	"fmt"
)

var m2 = map[string]myPackage.VertexMap{
  	"Bell Labs": {
    	40.68433, -74.39967,
  	},
  	"Google": {
    	37.42202, -122.08408,
  	},
}

func main() {
    fmt.Println(m2)
}
