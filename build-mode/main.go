package main

import "C"
import "fmt"

//export hello_from_lib
func hello_from_lib(a C.int) {
	fmt.Printf("Hello from Go, a = %d\n", a)
}

func main() {

}
