package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Newton's method to call sqrt of x
	//https://en.wikipedia.org/wiki/Newton%27s_method
	z := 1.0

	for i := 0; i <= 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}
	fmt.Println("Actual Square root", math.Sqrt(x))
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
