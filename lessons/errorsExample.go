package main

import (
    "fmt"
    "strconv"
)


func main() {
    convertStr("42")
    convertStr("abba")
}


func convertStr(str string) {
    i, err := strconv.Atoi(str)

    if err != nil {
        fmt.Printf("Convert error, %v\n", err)
        return
    }

    fmt.Println("Converted number:", i)
}
