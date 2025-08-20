package main

import "C"
import "fmt"

//export hello_from_lib1
func hello_from_lib1(name *C.char) {
	fmt.Printf("lib1: Hello %s!\n", C.GoString(name))
}

func main() {

}
