package main

import (
	"../goLang/validate"
)

type MyReader struct{}

func (rd MyReader) Read(s []byte) (int, error) {
	for i := 0; i < len(s); i++ {
		s[i] = 'A'
	}

	return len(s), nil
}

func main() {
	validate.Validate(MyReader{})
}
