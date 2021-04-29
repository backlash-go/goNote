package main

import "fmt"

func main() {

	//arr1 := []int{1,2,3}
	sum(1, 2, 3)
	nextNumFunc := nextNUm()

	fmt.Println(nextNumFunc())
	fmt.Println(nextNumFunc())
	fmt.Println(nextNumFunc())
	i := 100

	fmt.Println(fact(3))
	fmt.Println(i)
	test()
	test()
	test()

}

func sum(nums ...int) {
	fmt.Printf("%T\n", nums)
	fmt.Print(nums, " ") //输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}

func nextNUm() func() int {
	i, j := 1, 1
	fmt.Println(i,j)
	fmt.Println("aaaaa")


	return func() int {
		var tmp = i + j
		i, j = j, tmp
		fmt.Println(i,j)

		return tmp
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func test()  {
	i,j := 1,1
	var tmp = i + j
	i, j = j, tmp
	fmt.Println(i,j)

}