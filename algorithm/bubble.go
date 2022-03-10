package main

import "fmt"

const (
	LENGTH = 6
)

func main() {
	number := []int64{5, 6, 2, 3, 4, 1}
	bubbleSort(number)

}

func bubbleSort(nums []int64) {
	flags1 := false
	for i := 0; i < len(nums); i++ {

		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flags1 = true
			}
		}
		if !flags1 {
			break
		}
	}
	fmt.Println(nums)
}
