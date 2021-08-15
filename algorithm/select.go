package main

import "fmt"

func main() {
	number := []int{7, 6, 2, 3, 4, 1}
	selectionSort(number)
	fmt.Println(number)
}

/*

i= 0   min= 0    j=1
j := i + 1; j < length; j++
{
  arr[0] > arr[1]  min=j=1

  arr[1] > arr[2]  min=j=2

  arr[2] > arr[5]   min=j=5   		arr[i], arr[min] = arr[min], arr[i]    arr[0], arr[5] = arr[5], arr[0]

  {1, 6, 2, 3, 4, 7}
}



 */

func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
