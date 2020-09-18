package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 16, 16)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{100, 100, 255, 255}
}

func main() {
	m := Image{}

	// customized image
	pic.ShowImage(m)

	fmt.Printf("\n")

	// 256 * 256, purely blue image
	pic.Show(func(dx, dy int) [][]uint8 {
		points := make([][]uint8, 256)
		for i := range points {
			points[i] = make([]uint8, 256)
		}
		return points
	})
}
