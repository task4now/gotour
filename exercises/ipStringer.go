package main

import (
	"fmt"
	"strconv"
)

type IPaddr [4]byte

func (ip IPaddr) String() string {
	str := ""

	for i := 0; i < len(ip); i++ {
		str += strconv.Itoa(int(ip[i]))

		if i != len(ip)-1 {
			str += "."
		}
	}

	return str
}

func main() {
	hosts := map[string]IPaddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
