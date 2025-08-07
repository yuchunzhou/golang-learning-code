package main

import (
	"fmt"
	"time"
)

func panicFunc() {
	defer func() {
		r := recover()
		fmt.Printf("r in panicFunc func: %v\n", r)
	}()

	panic("oops")
}

func main() {
	r := recover()
	fmt.Printf("r in main func: %v\n", r)

	go func() {
		for {
			fmt.Println("another goroutine is runnning ...")
			time.Sleep(time.Second * 1)
		}
	}()

	panicFunc()
	time.Sleep(time.Second * 3)
}
