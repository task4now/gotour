package main

import (
	"../myPackage"
	"fmt"
	"strconv"
)

func splitNumberByDigits(number int) string {
	oldStr := strconv.Itoa(number)
	newStr := ""

	for i := 0; i < len(oldStr); i++ {
		if i != len(oldStr) -1 {
			newStr += string(oldStr[i]) + ", "
		} else {
			newStr += string(oldStr[i]) + "."
		}
	}

  return newStr
}

func main() {
	number := myPackage.GetIntFromInput(
		"Please, input number to split it on digits: ",
		9999,
		)

	digits := splitNumberByDigits(number)
	fmt.Print("Digits are: ", digits)
}
