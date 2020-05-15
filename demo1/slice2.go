package main

import "fmt"

func main() {


	/*
	切片是一种轻量级的数据结构，由指针，长度，容量组成，底层引用了一个数组对象，可以访问数组的子序列或者全部元素


	扩容机制，具体查看 runtime包下的 growslices 函数以及相关方法
	 */
	var slice1 = make([]int64, 5, 5)
	var finallyCap int
	var oldCap = cap(slice1) // 老容量
	var newCap = 2 * oldCap  // 2倍容量
	//if newCap > 1024 newCap = oldCap +oldCap/4
	var nowCap = cap(slice1) + len([]int64{1, 2, 3, 4}) // 新增之后一共的需要的长度
	// cap < 1024
	if newCap > nowCap {
		finallyCap = newCap
	} else {
		finallyCap = nowCap
		if finallyCap%2 != 0 {
			finallyCap = nowCap + 1
		}
	}
	fmt.Println(finallyCap)
	slice1 = append(slice1, []int64{1, 2, 3, 4, 5, 6}...)
	fmt.Println(len(slice1), cap(slice1))  //11 12
}
