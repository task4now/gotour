package main

import (
	"fmt"
	"strconv"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + strconv.FormatFloat(float64(e), 'f', 6, 64)
}

func SqrtNegError(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	value := float64(1)

	for i := 0; i < 10; i++ {
		value = 0.5 * (value + x/value)
	}

	return value, nil
}

func main() {
	fmt.Println(SqrtNegError(2))
	fmt.Println(SqrtNegError(-2))
}
