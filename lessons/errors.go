package main

import (
    "fmt"
    "time"
)

type MyError struct {
    When time.Time
    What string
}


func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}


func runError() error {
    return &MyError{
        time.Now(),
        "It didn't work...",
    }
}


func main() {
    if err := runError(); err != nil {
        fmt.Println(err)
    }
}
