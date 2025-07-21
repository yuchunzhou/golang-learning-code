package main

import (
	"fmt"
	"reflect"
	"sync"
)

func foo(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 1
}

func main() {
	var a map[string]int8
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a == nil)

	// a is nil, below code line will throw an error
	// a["a"] = 1

	a = make(map[string]int8)
	a["a"] = 1
	fmt.Println(a)

	var b chan int
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b == nil)

	// b is nil, below code block will throw an error
	// wg.Add(1)
	// go foo(b)
	// wg.Wait()

	var wg sync.WaitGroup

	b = make(chan int, 1)

	wg.Add(1)
	go foo(&wg, b)
	wg.Wait()

	fmt.Println(<-b)
	close(b)

	var c []string
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c == nil)

	// c is nil and has length 0, below code line will throw an error
	// fmt.Println(c[0])

	c = make([]string, 1)
	fmt.Println(c[0])
	c = append(c, "hello")
	fmt.Println(c)
}
