package main

import (
	"fmt"
	"time"
)

// variable scoping:
// defer is lexical binding
// goroutine is volatile
// closure is volatile

func foo() func() {
	a := 1
	defer fmt.Printf("defer 1: %d %p\n", a, &a) // 1

	go func() {
		fmt.Printf("goroutine 1: %d %p\n", a, &a)
	}()

	f := func() {
		fmt.Printf("in closure: %d %p\n", a, &a)
	}

	a++
	defer func() {
		fmt.Printf("defer 2: %d %p\n", a, &a) // 2
	}()

	go func() {
		fmt.Printf("goroutine 2: %d %p\n", a, &a)
	}()

	return f
}

func main() {
	f := foo()
	f() // 2

	time.Sleep(1 * time.Second)
}
