package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

// Pic shows me money
func Pic(dx, dy int) [][]uint8 {
	row := []uint8{}
	for i := 0; i < dy; i++ {
		row = append(row, uint8(dx))
	}

	rows := [][]uint8{}
	for i := 0; i < dy; i++ {
		rows = append(rows, row)
	}

	fmt.Println(rows)
	return rows
}

func main() {
	pic.Show(Pic)
}
