package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (int, float64) {
	z := 1.0
	loops := 0
	zPrev := 0.0

	for math.Abs(z-zPrev) > 0.00000001 {
		loops++
		zPrev = z
		z -= (z*z - x) / (2 * z)
	}
	return loops, z
}

func main() {
	num := 2.0

	loops, ans := Sqrt(num)
	fmt.Printf("Loops: %d\n", loops)
	fmt.Println(ans)
	fmt.Println(math.Sqrt(num))
}
