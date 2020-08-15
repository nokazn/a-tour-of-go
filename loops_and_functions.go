package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for diff := z; diff*diff > 1e-10; z -= diff {
		diff = (z*z - x) / (2 * z)
		fmt.Printf("The exponentiation of %v is %v\n", z, z*z)
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(2),
		math.Sqrt(2),
	)
}
