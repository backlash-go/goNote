package main

import "fmt"

func main() {

	ch := make(chan int)

	ch2 := make(chan<- int)

	ch3 := make(<-chan int)

	ch4 := make(chan string)

	go func() {
		ch <- 43

		ch2 <- 50

		ch4 <- "sadas"
	}()

	value := <-ch

	v3 := <-ch3

	fmt.Println(value)
	fmt.Println(v3)

}
