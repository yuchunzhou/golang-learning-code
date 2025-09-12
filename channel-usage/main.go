package main

import (
	"fmt"
	"time"
)

func main() {
	chan0 := make(chan any, 1)
	go func() {
		chan0 <- 1
	}()
	time.Sleep(1 * time.Second)
	fmt.Printf("ch0 output: %d\n", <-chan0)

	ch1 := make(chan int)
	go func() {
		fmt.Printf("ch1 output: %d\n", <-ch1) // will never output
		fmt.Println("111 ...")                // will never output
	}()

	time.Sleep(1 * time.Second)

	ch2 := make(chan int, 1)
	go func() {
		fmt.Printf("ch2 output: %d\n", <-ch2) // wiil never output
		fmt.Println("222 ...")                // will never output
	}()

	time.Sleep(1 * time.Second)

	ch3 := make(chan int)
	go func() {
		fmt.Printf("ch3-1 output: %d\n", <-ch3) // will output 1
	}()

	ch3 <- 1

	go func() {
		fmt.Printf("ch3-2 output: %d\n", <-ch3) // will output 2
	}()

	ch3 <- 2

	time.Sleep(1 * time.Second)

	ch4 := make(chan int, 3)
	go func() {
		fmt.Printf("ch4-1 output: %d\n", <-ch4)
	}()

	ch4 <- 3
	ch4 <- 4

	go func() {
		fmt.Printf("ch4-2 output: %d\n", <-ch4)
	}()

	time.Sleep(1 * time.Second)
}
