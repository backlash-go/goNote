package main

import "fmt"

func main() {
	//常量的定义    一般使用大写字母
	const A int =  10
	const B  =  20
	fmt.Printf("%T\n",A) //int
	fmt.Printf("%T\n",B) //int
	fmt.Println(A,B) //10 20

	const (
		D = 100
		E = 200
	)
	fmt.Println(D,E) //100 200


	//常量可以参与程序的运算
	c := A + 100
	fmt.Println(c) //110

	//常量不能修改
	//a = a + 10 //cannot assign to a
	//fmt.Println(a)

	//常量地址不允许访问，必须定义
	//fmt.Printf("%p\n",&D)  //cannot take the address of D
}
