package main

import (
	"fmt"
	"math"
)

func Sqrt(num float64) float64 {
	z := 1.0
	delta := 0.0
	for {
		z = z - (z*z-num)/(2*z)
		if math.Abs(delta-z) < 1e-10 {
			break
		}
		delta = z
	}
	return delta
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
