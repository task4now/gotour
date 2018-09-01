package main

import "fmt"

func main() {
  	names := [4]string{
    	"task4Now",
    	"Ridj",
    	"Kato Basiro",
    	"Mr.Dad",
  	}

  	fmt.Println(names)

  	a := names[0:2]
  	b := names[1:3]
  	fmt.Println(a, b)

  	b[0] = "Forever"
  	fmt.Println(a, b)
  	fmt.Println(names)
}
