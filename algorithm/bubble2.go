package main

import "fmt"

func main() {

	var num = []int{6, 5, 4, 1, 2, 3}
	fmt.Println(Bubble(num))

}

func Bubble(num []int) []int {

	flag := false

	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				flag = true
			}
		}

		if !flag {
			break

		}

	}

	return num

}
