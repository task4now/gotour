package main

import (
	"../myPackage"
	"fmt"
)

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	Truth := myPackage.GetRandomInt(10) > 0

	fmt.Println("Go rules?", Truth)
}
