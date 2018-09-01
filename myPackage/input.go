package myPackage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func GetIntFromInput(intro string, seed int) int {
	var number int

	str := GetStrFromInput(intro, "")

	if str == "" {
		number = GetRandomInt(seed)
		fmt.Println("Will be used random number from seed:", number)
	} else {
		var err error
		number, err = strconv.Atoi(str)

		if err != nil {
			number = GetRandomInt(seed)
			fmt.Println("Your input is not a number, so, will be used random number from seed:", number)
		}
	}

	return number
}


func GetStrFromInput(intro string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(intro)
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("There was occured error:", err)
		return defaultValue
	}

	if len(str) < 2 {
		fmt.Println("You don't input anything.")
		return defaultValue
	}

	return str[:len(str)-1]
}
