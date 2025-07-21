package main

import "fmt"

type Foo struct {
	data string
}

func test_struct(input Foo) Foo {
	input.data = "!!!"
	return input
}

func test_slice(input []string) []string {
	input[0] = "!!!"
	return input
}

func test_map(input map[string]int) map[string]int {
	input["hello"] = 1
	return input
}

func test_array(input [1]string) [1]string {
	input[0] = "!!!"
	return input
}

func test_string(input string) string {
	input = "!!!"
	return input
}

func test_chan(input chan int) chan int {
	input <- 1
	return input
}

func main() {
	// struct is copy action
	a := Foo{
		data: "hello",
	}
	b := test_struct(a)
	fmt.Println(a) // hello
	fmt.Println(b) // !!!

	// slice is reference action
	e := []string{"hello"}
	f := test_slice(e)
	fmt.Println(e) // !!!
	fmt.Println(f) // !!!

	// map is reference action
	g := map[string]int{
		"hello": 0,
	}
	h := test_map(g)
	fmt.Println(g) // hello 1
	fmt.Println(h) // hello 1

	// array is copy action
	i := [1]string{"hello"}
	j := test_array(i)
	fmt.Println(i) // hello
	fmt.Println(j) // !!!

	// string is copy action
	k := "hello"
	l := test_string(k)
	fmt.Println(k) // hello
	fmt.Println(l) // !!!

	// chan is reference action
	m := make(chan int, 2)
	m <- -1
	_ = test_chan(m)
	fmt.Println(<-m) // -1
	fmt.Println(<-m) // 1
}
