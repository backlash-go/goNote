package main

import "fmt"

func main() {
	fmt.Println("Start")
	doSomething()
	fmt.Println("End")
}



func doSomething() {
	defer fmt.Println("Deferred call in doSomething")
	fmt.Println("Start doSomething")
	panic("something went wrong") // 引发 panic
	fmt.Println("End doSomething") // 不会执行到这里
}