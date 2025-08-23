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

type Foo struct {
	value int
}

// foo is just a copy
func (foo Foo) SayHello() {
	fmt.Println("hello from Foo")
	foo.value = 1
}

type Bar struct{}

func (bar *Bar) SayHello() {
	fmt.Println("hello from Bar")
}

func main() {
	foo := Foo{}
	helloFunc(foo)
	fmt.Printf("%d\n", foo.value) // 0

	bar := Bar{}
	helloFunc(&bar)
}
