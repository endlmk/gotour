package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	Walk1(t, ch)
	close(ch)
}
func Walk1(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk1(t.Left, ch)
	} 
	ch <- t.Value
	if t.Right != nil {
		Walk1(t.Right, ch)
	} 
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for v1 := range c1{
		v2 := <- c2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int, 10)
	go Walk(tree.New(1), c)
	for i := range c {
		fmt.Println(i)
	}
	r1 := Same(tree.New(1), tree.New(1))
	fmt.Println(r1)
	r2 := Same(tree.New(1), tree.New(2))
	fmt.Println(r2)
}