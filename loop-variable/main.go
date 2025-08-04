package main

import (
	"fmt"
	"time"
)

func foo() func() {
	data_ := "!!!"
	fmt.Printf("in foo: address of data_ is %p\n", &data_)

	return func() {
		fmt.Printf("in closure: address of data_ is %p, value is %s\n", &data_, data_) // data_ escapes to heap
	} // function also escapes to heap
}

func main() {
	data := "hello" // data escapes to heap
	go func() {
		fmt.Printf("line 19: address is %p, value is %s\n", &data, data)
	}() // function escapes to heap

	go func() {
		fmt.Printf("line 24: address is %p, value is %s\n", &data, data)
	}()

	data = "world"              // main goroutine is still running
	time.Sleep(1 * time.Second) // pause the main goroutine
	data = "WORLD"

	time.Sleep(3 * time.Second)

	f := foo()
	f()

	time.Sleep(3 * time.Second)

	values := []string{"aaa", "bbb", "ccc"}
	for _, value := range values { // value escapes to heap
		fmt.Printf("address of value: %p\n", &value) // three value variables have different addresses
		go func() {
			fmt.Println(value)
		}()
	}

	time.Sleep(3 * time.Second)
}

// go build -gcflags "-m" main.go
