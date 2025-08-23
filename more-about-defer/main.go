package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func foo(a int) func() {
	// defer statement
	// arguments evaluated immediately
	defer fmt.Printf("defer 1: %d %p\n", a, &a) // 1

	// defer with function literal
	// closures capture variables by reference
	defer func() {
		fmt.Printf("defer 2: %d %p\n", a, &a) // 2
	}()

	go func() {
		fmt.Printf("goroutine 1: %d %p\n", a, &a) // 2
	}()

	f := func() {
		fmt.Printf("in closure: %d %p\n", a, &a) // 2
	}

	a++
	defer func() {
		fmt.Printf("defer 3: %d %p\n", a, &a) // 2
	}()

	go func() {
		fmt.Printf("goroutine 2: %d %p\n", a, &a) // 2
	}()

	return f
}

func main() {
	var arg int
	arg, _ = strconv.Atoi(os.Args[1])

	f := foo(arg)
	f()

	time.Sleep(1 * time.Second)
}
