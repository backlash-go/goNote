package main

import (
	"fmt"
	"time"
)

/*

只有读channel  会出现死锁，
主GO程自己 写channel 读channel 会出现死锁


1 必须有一个写端，与 读端，

2  读写两端对应的应该是不同的GO程， 不能在同一个GO 程中自己读自己写
*/

func main() {
	chan1 := make(chan string)

	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("子GO 程 i=", i)
			time.Sleep(time.Second)
		}
		chan1 <- "abcd"
	}()
	str := <-chan1
	fmt.Println("主程，", str)
}

func main002() {
	//创建一个用于通信的Channel
	chan1 := make(chan string)

	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("子GO 程 i=", i)
			time.Sleep(time.Second)
		}
	}()

	//主GO程自己 写channel 读channel 会出现死锁
	//fatal error: all goroutines are asleep - deadlock!

	chan1 <- "abcd"

	str := <-chan1
	fmt.Println("主程，", str)
}

func main001() {
	//创建一个用于通信的Channel
	chan1 := make(chan string)
	//只有读channel  会出现死锁，

	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("子GO 程 i=", i)
			time.Sleep(time.Second)
		}
	}()
	//fatal error: all goroutines are asleep - deadlock!

	str := <-chan1
	fmt.Println("主程，", str)
}
