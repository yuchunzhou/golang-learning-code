package main

import (
	"fmt"
	"reflect"
)

type integer interface {
	int | int8 | int16 | int32 | int64
}

func foo[E integer](n E) {
	fmt.Println(reflect.TypeOf(n))
}

func fuzz[E ~string](s []E) {
	fmt.Println(len(s))
}

func main() {
	foo(int16(1))
	fuzz([]string{"hello", "world"})
}
