package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{q[0], r[0]},
		{q[1], r[1]},
		{q[2], r[2]},
		{q[3], r[3]},
		{q[4], r[4]},
		{q[5], r[5]},
	}

	fmt.Println(s)
}
