package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func task_func(ctx context.Context, args string, sync_ch chan int) {
	fmt.Println("task is running ...")

	for {
		select {
		case <-ctx.Done():
			// context is canceled or timeout
			fmt.Printf("context is canceled or timeout: %s, args: %s\n", ctx.Err().Error(), args)
			return
		case <-sync_ch:
			fmt.Printf("task should exit now, args: %s\n", args)
			return
		default:
			fmt.Println("task is still running ...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// *context.timerCtx
	fmt.Printf("ctx type: %s\n", reflect.TypeOf(ctx))

	// buffered channel will not block
	sync_ch := make(chan int, 1)

	go task_func(ctx, "some args", sync_ch)

	time.Sleep(10 * time.Second)
	// cancel()
	sync_ch <- 1
	time.Sleep(1 * time.Second)
}
