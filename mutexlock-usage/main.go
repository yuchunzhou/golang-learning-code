package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var count int

	cpu_count := runtime.NumCPU()
	wg.Add(cpu_count)

	for i := 0; i < cpu_count; i++ {
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			count++
			mutex.Unlock()

			fmt.Printf("goroutine %d adds one\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Printf("last value of count: %d\n", cpu_count)
}
