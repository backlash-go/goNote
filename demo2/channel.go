package main

import "fmt"

func main() {
	ch1 := make(chan int,3)
	fmt.Println(ch1 == nil)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	//ch1 <- 4
	elem , ok:= <- ch1
	fmt.Println(elem,ok)
	elem = <- ch1
	fmt.Println(elem)


	//单向通道最主要的用途就是约束其他代码的行为
	//单向通道 只能发送，不能接收 ：简称为发送通道
	//sendSingle := make(chan <- int,1)
	//ch2 := <- sendSingle  //<-sendSingle (receive from send-only type chan<- int)


	//只能接收不能发送  ： 简称为接收通道
	//receive := make(<- chan int,2)
	//receive <- 2  // receive <- 2 (send to receive-only type <-chan int)

}
