package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	n := len(b)

	// make every element be 'A'
	for key := range b {
		b[key] = 'A'
	}

	return n, nil
}

func main() {
	reader.Validate(MyReader{})
}
