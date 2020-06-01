package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义互斥锁
var mutex sync.Mutex // 有空间的

func printer(str string) {
	mutex.Lock() // 对公共区加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock() // 公共区使用结束，解锁
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
