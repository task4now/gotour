package main

import (
	"../myPackage"
	"fmt"
)

func main() {
	// Random int realisation is in myRandom package
	fmt.Println("Today's random number is", myPackage.GetRandomInt(367))
}
