package main

import (
	"../myPackage"
	"fmt"
	"time"
)

func main() {
	name := myPackage.GetStrFromInput(
		"Hello, friend! What is your name?\n",
		"Anonymous")

    now := time.Now()

    year, month, day := now.Date()
	weekday := now.Weekday()
    hour, min, sec := now.Clock()

    secZero := ""
    if sec < 10 {
        secZero += "0"
    }

    fmt.Printf("\n\tSo, %s, it's nice to meet you!\n", name)
    fmt.Printf("\tToday is %s -  %d of %s, %d.\n", weekday, day, month, year)
    fmt.Printf("\tCurrent time is %d:%d:%s%d\n\n", hour, min, secZero, sec)
}
