package main

import "fmt"

func bar() int {
	val := 1

	defer func() {
		val += 1
	}()

	return val
}

type Foo struct {
	data string
}

// data := val
// val += 1
// return
func buz() (data int) {
	val := 1

	a := 1
	b := a // copied from a
	a = 2
	fmt.Printf("b = %d\n", b) // 1

	c := Foo{
		data: "hello",
	}
	d := c // copied from c
	c.data = "world"
	fmt.Printf("d = %#v\n", d) // hello

	defer func() {
		val += 1
	}()

	return val // val is still on stack, its value is copied to data
}

// data := 0
// data = data + 1
// return
func fuzz() (data int) {
	defer func() {
		fmt.Println("fuzz defer 1") // after defer 2, but before return
		data = data + 1
	}()

	return func() int {
		fmt.Println("fuzz defer 2") // before return
		return 0
	}()
}

func foo() func() int {
	defer fmt.Println("foo defer 1")
	defer fmt.Println("foo defer 2")

	data := 1

	defer func() {
		fmt.Println("foo defer 3")
		data++
	}()
	defer func(data int) {
		data = data + 1
		fmt.Printf("foo defer 4: %d\n", data)
	}(data)

	fmt.Println("in foo ...")

	bar := func() int {
		defer fmt.Println("bar defer 1")

		fmt.Println("in bar ...")
		return data // data is moved to heap
	}
	return bar
}

func main() {
	fmt.Println(foo()()) // 2
	fmt.Println(fuzz())  // 1
	fmt.Println(buz())   // 1
	fmt.Println(bar())   // 1
}
