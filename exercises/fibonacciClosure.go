package main

import "fmt"


func fibonacci() func() int {
    fib1 := 0
    fib2 := 0

    return func() int {
        if fib2 == 0 {
            fib2++
      		return fib1
    	}

    	fib1, fib2 = fib2, fib1 + fib2

    	return fib1
  	}
}


func main() {
  	f := fibonacci()

  	for i := 0; i < 10; i++ {
    	fmt.Println(f())
  	}
}
