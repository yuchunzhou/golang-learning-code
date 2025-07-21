package main

import "fmt"

type helloInterface interface {
	sayHello()
}

func helloFunc(who helloInterface) {
	who.sayHello()
}

type Foo struct{}

func (foo Foo) sayHello() {
	fmt.Println("hello from Foo")
}

type Bar struct{}

func (bar *Bar) sayHello() {
	fmt.Println("hello from Bar")
}

func main() {
	foo := Foo{}
	helloFunc(foo)

	bar := Bar{}
	helloFunc(&bar)
}
