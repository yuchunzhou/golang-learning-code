package main

import (
	"fmt"
	"reflect"
)

type HelloInterface interface {
	SayHello()
}

func helloFunc(who HelloInterface) {
	fmt.Printf("who's type: %s\n", reflect.TypeOf(who))
	who.SayHello()
}

type Foo struct{}

func (foo Foo) SayHello() {
	fmt.Println("hello from Foo")
}

type Bar struct{}

func (bar *Bar) SayHello() {
	fmt.Println("hello from Bar")
}

func main() {
	foo := Foo{}
	helloFunc(foo)

	bar := Bar{}
	helloFunc(&bar)

	// function name is a pointer
	fmt.Printf(">>> %p\n", helloFunc)
}
