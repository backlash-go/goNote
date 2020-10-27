package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*

sync包的waitgroup类型，它比通道更加适合实现这种一对多的goroutine协作流程
*/

type sss struct {
	slc [3]int64
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	sss := sss{}



	fmt.Printf("%T,%v,%v\n",sss.slc,len(sss.slc),cap(sss.slc))

	go func() {
		fmt.Println("8888")
		wg.Done()
	}()
	go func() {
		fmt.Printf("%T\n","a")
		wg.Done()
	}()
	wg.Wait()
}

func main0001() {

	sign := make(chan struct{})
	go func() {
		fmt.Println("hello")
		sign <- struct{}{}
	}()
	go func() {
		fmt.Println("baby")
		sign <- struct{}{}
	}()
	<-sign
	<-sign

}

func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		} else {
			//fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}
