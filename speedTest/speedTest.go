package main

import "fmt"

func main() {
  	sum := 1

  	for {
    	sum += 1

    	if sum == 10000000000 {
      		break
    	}
  	}

  	fmt.Println(sum)
}
