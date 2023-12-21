package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	fmt.Println(ch1 == nil)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	//ch1 <- 4
	elem, ok := <-ch1
	fmt.Println(elem, ok)
	elem = <-ch1
	fmt.Println(elem)


	receive := make(chan<- int, 2)
	receive <- 2 // receive <- 2 (send to receive-only type <-chan int)

}
