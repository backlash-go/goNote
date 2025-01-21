package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second) // 创建一个每秒触发一次的计时器

	for {
		select {

		case <-ticker.C:
			fmt.Println("Tick at")
		default:
			fmt.Println("sad")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
