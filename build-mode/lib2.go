package main

import "C"
import "fmt"

//export hello_from_lib2
func hello_from_lib2(name *C.char) {
	fmt.Printf("lib2: Hello %s!\n", C.GoString(name))
}

func main() {

}
