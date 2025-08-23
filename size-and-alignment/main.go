package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	// size: 4
	a int8 // align: 1 size: 1
	// padding 1 byte
	b int16 // align: 2 size: 2

	// size: 8
	c int8
	// padding 3 bytes
	d int32 // align: 4 size: 4

	// size: 12
	e int8
	// padding 3 bytes
	f int64 // align: 8 size: 8
}

type Bar struct {
	a int8
	// padding 7 bytes
	b Foo // align: 8 size: 24
}

func main() {
	foo := Foo{}
	// size: (1 + 1 + 2) + (1 + 3 + 4) + (1 + 3 + 8) = 24
	fmt.Printf("size: %d\n", unsafe.Sizeof(foo))
	// align: the max align bytes among all struct members
	fmt.Printf("align: %d\n", unsafe.Alignof(foo))

	bar := Bar{}
	// size: 1 + 7 + 24 = 32
	fmt.Printf("size: %d\n", unsafe.Sizeof(bar))
	// align: 8
	fmt.Printf("align: %d\n", unsafe.Alignof(bar))
}
