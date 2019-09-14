package main

import (
	"fmt"
)

func fibonacci() func() int {
	pp := 0
	p := 1
	return func() int {
		v := pp
		pp, p = p, v + p
		return v
	}
} 

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}