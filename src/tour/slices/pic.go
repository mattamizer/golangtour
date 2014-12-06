package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for ele := range slice {
		slice[ele] = make([]uint8, dx)
		for val := range slice[ele] {
			slice[ele][val] = uint8((dx ^ dy) * (dx ^ dy))
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
