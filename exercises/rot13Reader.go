package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)

	for i := 0; i < len(b); i++ {
		char := b[i]

		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			b[i] += 13
		} else if (char > 'M' && char <= 'Z') || (char > 'm' && char <= 'z') {
			b[i] -= 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
