package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	Num float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.Num)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt{x}
	}

	z := 1.0
	loops := 0
	zPrev := 0.0

	for math.Abs(z-zPrev) > 0.00000001 {
		loops++
		zPrev = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
