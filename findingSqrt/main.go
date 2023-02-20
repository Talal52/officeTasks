package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i < 10; i++ {
		z = z - (z*z-x)/(2*z)

	}
	return z
}

func main() {
	t := 25.0
	fmt.Println(Sqrt(t))
	fmt.Println(math.Sqrt(t))

}
