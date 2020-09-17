package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (z *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = z.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = convertRot13(p[i])
	}

	return n, err
}

func convertRot13(b byte) byte {
	var lowerBound, upperBound byte
	switch {
	case 'a' <= b && b <= 'z':
		lowerBound, upperBound = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		lowerBound, upperBound = 'A', 'Z'
	default:
		return b
	}

	if upperBound <= (b + 13) {
		return lowerBound + (b + 12 - upperBound)
	}
	return b + 13
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
