package main

import "fmt"

func main() {
	var a *int
	if a == nil {
		fmt.Println("a is nil")
	}

	b := new(int)

	if b == nil {
		fmt.Println("b is nil")
	}
	fmt.Println(*b)

}
