package main

import (
    "../goLang/tree"
    "fmt"
)


func Walk(t *tree.Tree, ch chan int) {

    ch <- t.Value

    if t.Left != nil {
        go Walk(t.Left, ch)
    }

    if t.Right != nil {
        go Walk(t.Right, ch)
    }
}


func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    var ar1 [10]int
    var ar2 [10]int

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for i := 0; i < 10; i++ {
        ar1[i] = <-ch1
        ar2[i] = <-ch2
    }


    for i := 0; i < 10; i++ {
        found := false

        for j := 0; j < 10; j++ {
            if ar1[i] == ar2[j] {
                found = true
                break
            }
        }

        if !found {
            return false
        }
    }

    return true
}


func main() {
    ch := make(chan int)
    go Walk(tree.New(2), ch)

    for i := 0; i < 10; i++ {
        v := <-ch
        fmt.Println(v)
    }

    fmt.Println(tree.New(2))

    fmt.Println(Same(tree.New(1), tree.New(2)))
    fmt.Println(Same(tree.New(3), tree.New(3)))
}
