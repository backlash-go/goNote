package main

import "fmt"

func main() {
	fmt.Println("return is ",fun2(1))
	fmt.Println("return is ",b())

	c:=c()
	fmt.Println("c return:", *c, c) // 打印结果为 c return: 2 0xc082008340
}


func fun2(i int) int {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
		fmt.Printf("%p,%s\n",&i,"defer1")

	}()
	fmt.Printf("%p,%s\n",&i,"de")

	return i
}


func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}


func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i, &i) // 打印结果为 c defer2: 2 0xc082008340
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i, &i) // 打印结果为 c defer1: 1 0xc082008340
	}()
	return &i
}

/*
多个defer的执行顺序为“后进先出”；

所有函数在执行RET返回指令之前，都会先检查是否存在defer语句，若存在则先逆序调用defer语句进行收尾工作再退出返回；

匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值；

return其实应该包含前后两个步骤：第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值）；第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，最后RET携带返回值退出函数；

‍‍因此，‍‍defer、return、返回值三者的执行顺序应该是：return最先给返回值赋值；接着defer开始执行一些收尾工作；最后RET指令携带返回值退出函数。
 */