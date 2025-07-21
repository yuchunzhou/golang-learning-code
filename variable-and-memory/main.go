package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// string
	a := "hello world"
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&a)))

	// number
	b := 1
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(reflect.TypeOf(b))

	// slice
	var c []int64 = make([]int64, 5)
	c[0] = 1
	c[1] = 2
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&c)))
	fmt.Println(reflect.TypeOf(c))

	// array
	d := [3]int64{1, 2, 3}
	fmt.Printf("%d\n", unsafe.Pointer(&d))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(reflect.TypeOf(d))

	// slice again
	e := d[:]
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&e)))
}
