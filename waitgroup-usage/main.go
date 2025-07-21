package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cpu_count := runtime.NumCPU()

	for i := 0; i < cpu_count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("this is goroutine " + strconv.Itoa(i))
		}(i)
	}

	wg.Wait()
}
