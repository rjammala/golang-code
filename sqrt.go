package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	old := 10.0
	const delta = 1e-6
	for {
		old = z - (math.Pow(z, 2)-x)/(2*z)
		if math.Abs(old-z) < delta {
			break
		}
		z = old
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
