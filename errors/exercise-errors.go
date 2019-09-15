package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	num := fmt.Sprint(float64(e))
	return "cannot Sqrt negative number: " + num
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} 
	z := x / 2
	i := 0
	const delta = 1e-9
	for {
		diff := (z * z - x)/(2 * x)
		if math.Abs(diff) < delta {
			return z, nil
		} 
		z -= diff
		i++
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}