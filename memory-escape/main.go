package main

import "fmt"

type Foo struct {
	data string
}

func foo() *Foo {
	foo_ := Foo{
		data: "hello",
	}

	fmt.Printf("%p\n", &foo_)
	return &foo_
}

func main() {
	foo_ := foo()
	fmt.Printf("%p\n", foo_)
	fmt.Println(foo_.data)
}

// go build -gcflags "-m" main.go
