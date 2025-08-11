package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make(chan int, 10)

	go func() {
		for i := 0; i < 5; i++ {
			numbers <- i
			time.Sleep(1 * time.Second)
		}

		close(numbers)
	}()

	for number := range numbers {
		fmt.Println(number)
	}

	fmt.Println("Channel is closed")
}
