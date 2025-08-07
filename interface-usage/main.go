package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type HelloInterface interface {
	SayHello()
}

func helloFunc(who HelloInterface) {
	fmt.Printf("who ptr: %p\n", &who)

	// do reflection
	who_type := reflect.TypeOf(who)
	if who_type.Kind() == reflect.Ptr {
		fmt.Printf("who's type name: *%s\n", who_type.Elem().Name())
	} else {
		fmt.Printf("who's type name: %s\n", who_type.Name())
	}

	who.SayHello()
}

type Foo struct{}

func (foo Foo) SayHello() {
	fmt.Println("hello from Foo")
	fmt.Printf("foo ptr: %p\n", &foo)
}

type Bar struct{}

func (bar *Bar) SayHello() {
	fmt.Println("hello from Bar")
}

func main() {
	foo := Foo{}
	// do reflection
	foo_type := reflect.TypeOf(foo)
	fmt.Println(foo_type.NumMethod()) // get exported methods’ count
	for index := range foo_type.NumMethod() {
		fmt.Printf("method name %d of %s: %s, address: %p\n", index, foo_type.Name(), foo_type.Method(index).Name, unsafe.Pointer(foo_type.Method(index).Func.Pointer()))
	}

	helloFunc(foo)
	fmt.Printf("foo ptr: %p\n", &foo)

	bar := Bar{}
	helloFunc(&bar)

	// function name is a pointer
	fmt.Printf(">>> %p\n", helloFunc)
}
