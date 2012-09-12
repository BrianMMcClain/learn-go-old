package main

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dxy := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		dxy[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			dxy[x][y] = uint8((x^y))
		}
	}
	
	return dxy
}

func main() {
	pic.Show(Pic)
}