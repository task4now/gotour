package main

import "fmt"

type I4 interface {
    M()
}

func main() {
    var i I4
    describe3(i)
    i.M()
}

func describe3(i I4) {
    fmt.Printf("(%v, %T)\n", i, i)
}
