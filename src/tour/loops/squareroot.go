package main

import (
	"fmt"
)

func Sqrt(num float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z = z - ((z*z - num) / 2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
