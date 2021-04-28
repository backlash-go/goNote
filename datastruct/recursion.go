package main

import "fmt"

func f(n int) int {
	if n == 1 {
		return 1
	}
	return f(n-1) + 1
}



func main() {
	fmt.Println(f(9))
}
