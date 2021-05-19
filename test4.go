package main

import "fmt"

func main() {


	var ar1 [3]string
	ar1[2] = "100"
	fmt.Println(ar1)
}
func d1(n int) int {
	var sum int
	if n == 1 {
		sum = 1
		return sum
	}
	sum = n + d1(n-1)
	return sum

}


