package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (z float64, iter int) {
	z = x / 2
	i := 0
	const delta = 1e-9
	for {
		diff := (z * z - x)/(2 * x)
		if math.Abs(diff) < delta {
			iter = i
			return z, iter
		} 
		z -= diff
		i++
	}
}

func main() {
	fmt.Println(sqrt(1000))
}