package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// variable has its own memory
	// assgnment just modifies the memory content
	a := 1
	fmt.Printf("a address 1: %p\n", &a)

	a = 2
	fmt.Printf("a address 2: %p\n", &a)

	// make just for ref types: slice map chan
	// m is a pointer, points to the underlying dada struct
	m := make(map[string]int, 5)
	m["a"] = 1
	fmt.Printf("m address 1: %p %T %d\n", &m, m, unsafe.Sizeof(m))

	// previous map is GCed

	m = make(map[string]int, 10)
	m["a"] = 2
	m["b"] = 3
	fmt.Printf("m address 2: %p %T %d\n", &m, m, unsafe.Sizeof(m))

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
