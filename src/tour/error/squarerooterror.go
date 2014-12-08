package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt a negative number: %g", e)
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegativeSqrt(num)
	}
	z := 1.0
	delta := 0.0
	for {
		z = z - (z*z-num)/(2*z)
		if math.Abs(delta-z) < 1e-10 {
			break
		}
		delta = z
	}
	return delta, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(-2))
}
