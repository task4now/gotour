package main

import (
    "../myPackage"
)


func main() {
    var i interface{}
    myPackage.DescribeInterface(i)

    i = 42
    myPackage.DescribeInterface(i)

    i = "Hello!"
    myPackage.DescribeInterface(i)
}
