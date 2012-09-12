package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (((z*z)-x) / (2 * z))	
	}
	
	return z
}

func main() {
	z := float64(4)

	fmt.Println(Sqrt(z))
	fmt.Println(math.Sqrt(z))
}