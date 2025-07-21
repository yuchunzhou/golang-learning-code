package main

import (
	"context"
	"fmt"
	"time"
)

func task_func(ctx context.Context, args string, sync_ch chan int) {
	fmt.Println("task is running ...")

	// retrieve value from context
	ctx_value := ctx.Value(Key("aaa")).(*Value)
	fmt.Printf("context value: %v and value addr: %p\n", *ctx_value, ctx_value)

	for {
		select {
		case <-ctx.Done():
			// context is canceled or reached the deadline
			fmt.Printf("context is canceled or reached the deadline: %s, args: %s\n", ctx.Err().Error(), args)
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

type Key string
type Value int

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	// store a value in context
	value := Value(1)
	fmt.Printf("value addr: %p\n", &value)
	ctx = context.WithValue(ctx, Key("aaa"), &value)

	// write data into buffered channel will not block when it has capacity
	sync_ch := make(chan int, 1)

	go task_func(ctx, "some args", sync_ch)

	time.Sleep(3 * time.Second)
	// cancel()
	sync_ch <- 1
	time.Sleep(2 * time.Second)
}
