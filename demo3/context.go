package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func main002() {
	c := make(chan struct{})
	bv := make(chan struct{})

	go func() {
		fmt.Println(2222)
		close(c)
	}()

	go func() {
		fmt.Println(333)
		close(bv)
	}()

	<-c
	<-bv
}

func main001() {
	//var total, num int32= 12, 12
	total, num := 12, 11

	fmt.Printf("The number: %d [with context.Context]\n", num)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("555")
		cancel()
	}()

	<-ctx.Done()

	fmt.Println(total)

}
