package main

import "fmt"

func foo() (value int) {
	value = 1

	defer func() {
		value += 1
	}()

	return
}

func bar() (value int) {
	defer func() {
		value += 1
	}()

	return 1
}

func main() {
	fmt.Println(foo()) // 2
	fmt.Println(bar()) // 2
}
