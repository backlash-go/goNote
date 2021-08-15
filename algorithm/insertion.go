package main

import "fmt"

func main() {
	number := []int{5, 6, 2, 3, 4, 1}
	insertSort(number)
	fmt.Println(number)
}

/*
i = 1    j=0   value = arr[1] = 6
i = 2    j=1  value = arr[2] = 2    {2, 5, 6, 3, 4, 1}
i = 3    j=2   value = arr[3] = 3     {2, 3, 5, 6, 4, 1}


*/

func insertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}
