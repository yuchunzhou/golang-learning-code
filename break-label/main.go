package main

import "fmt"

func main() {
out_of_loop:
	for {
		fmt.Println("loop ...")
		break out_of_loop
	}
}
