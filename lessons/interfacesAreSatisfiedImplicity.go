package main

import "fmt"

type I1 interface {
  	M()
}

type T1 struct {
  	S string
}


func (t T1) M() {
  	fmt.Println(t.S)
}


func main() {
  	var i I1 = T1{"Hello!"}
  	i.M()
}
