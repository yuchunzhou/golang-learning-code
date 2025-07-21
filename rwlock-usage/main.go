package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	var count int

	go func() {
		for {
			mutex.RLock()
			fmt.Printf("count is %d\n", count)
			mutex.RUnlock()
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	cpu_count := runtime.NumCPU()
	wg.Add(cpu_count)

	for i := 0; i < cpu_count; i++ {
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			fmt.Printf("goroutine %d adds one\n", i)
			count++
			mutex.Unlock()
		}(i)
	}

	time.Sleep(time.Millisecond * 5000)
	wg.Wait()
}
