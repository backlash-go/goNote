package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string) // 带缓冲的通道
	ch1 <- "asdsad"
	//go func() {
	//	ch1 <- "asdsad"
	//}()

}

func s(ch <-chan string) {
	for {
		select {
		case msg := <-ch:
			fmt.Println("Received:", msg)
			time.Sleep(100 * time.Millisecond) // 避免忙等

		default:
			fmt.Println("default")
			time.Sleep(100 * time.Millisecond) // 避免忙等
		}
	}
}

