package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func goroutine_id() int {
	var buf = make([]byte, 100)
	var stack = buf[:runtime.Stack(buf, false)]
	id := strings.Split(string(stack), " ")[1]
	id_, _ := strconv.Atoi(id)
	return id_
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	for _ = range 1000 {
		go func() {
			for {
			}
		}()
	}

	for {
		fmt.Printf("there are %d goroutines running, current goroutine id is %d\n", runtime.NumGoroutine(), goroutine_id())
		runtime.Gosched()
	}
}
