package main

import (
	"fmt"
	"iter"
)

func forward[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func fib(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	slice := []string{"hello", "world"}
	for i, e := range forward(slice) {
		fmt.Printf("slice[%d] = %s\n", i, e)
	}

	cnt := 0
	for n := range fib(10) {
		cnt++
		fmt.Printf("%d ", n)
		if cnt == 10 {
			break
		}
	}
	fmt.Printf("\n")
}
