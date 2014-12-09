package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
)

type Image [][]uint8

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	dx, dy := len(i), 0
	if dx > 0 {
		dy = len(i[0])
	}
	return image.Rect(0, 0, dx, dy)
}

func (i Image) At(x, y int) color.Color {
	val := i[x][y]
	return color.RGBA{val, val, 255, 255}
}

func MakeImage(dx, dy int, f func(int, int) uint8) Image {
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
	m := MakeImage(255, 255, func(a, b int) uint8 { return uint8(a * b) })
	pic.ShowImage(m)
}
