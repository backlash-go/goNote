package main

import (
	"fmt"
	"time"
)


var channel1 = make(chan int)

func print1 (s string){
	for _, v := range s{
		fmt.Printf("%c",v)
		time.Sleep(5000000000)
	}
	channel1 <- 10
}


func print2 (s string){
	<- channel1
	for _, v := range s{
		fmt.Printf("%c",v)
		time.Sleep(5000000000)
	}
}
func main() {


	go print1("hello")
	go print2("world")

	for{
		;
	}

}
