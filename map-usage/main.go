package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	v, ok := m["a"]
	fmt.Println(v)
	fmt.Println(ok)

	v, ok = m["d"]
	fmt.Println(v)
	fmt.Println(ok)
}
