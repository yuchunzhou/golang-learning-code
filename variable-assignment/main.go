package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// local variables are allocated on stack
// assignment to variables just modifies the stack memory content

func main() {
	a := 1
	fmt.Printf("a(%T) address 1: %p\n", &a, &a)

	a = 2
	fmt.Printf("a(%T) address 2: %p\n", &a, &a)

	// make just for ref types: slice map chan
	m := make(map[string]int, 5)
	m["a"] = 1
	fmt.Printf("m(%T) address 1: %p %T %d\n", &m, &m, m, unsafe.Sizeof(m))

	runtime.GC()

	m = make(map[string]int, 100)
	m["a"] = 2
	m["b"] = 3
	fmt.Printf("m(%T) address 2: %p %T %d\n", &m, &m, m, unsafe.Sizeof(m))

	// new returns a pointer
	x := new(int)
	fmt.Printf("x: %p %d\n", x, *x)
	*x = 2
	fmt.Printf("x: %p %d\n", x, *x)

	// new is not for ref types
	y := new(map[string]int) // y is a nil pointer
	// (*y)["a"] = 1            // will panic
	fmt.Printf("y: %p %v\n", y, y)

	z := new(chan int) // nil channel
	fmt.Printf("z: %p %v\n", z, z)
	go func() {
		(*z) <- 1 // will panic
	}()
	fmt.Println(<-(*z)) // will panic
}
