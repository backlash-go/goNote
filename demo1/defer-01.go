package main

import "fmt"

func testDefer() int {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("returning 3")
	return 3
}

func main() {
	result := testDefer()
	fmt.Println("result:", result)
}
