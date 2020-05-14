package main

import (
	"fmt"
	"math"
)

func main01() {
	// var 变量名 数据类型 = 值
	var a1 int       //变量声明  没有赋值默认为0
	a1 = 10

	//a1 = a1 + 10      //表达式
	var b1 int =  20 //变量声明并赋值  即变量的定义
	fmt.Printf("%v\n%v\n",a1,b1)

	var a2 float64 = 2
	a2 = math.Pow(a2,2.5)
	fmt.Println(a2)

	// 自动推导类型
	a3 := "xxb"
	fmt.Println(a3)

	//不同的数据类型不能计算

	//交换两个变量的值

	var b2 int = 20

	a1 = a1 + b2
	b2 = a1 - b2
	a1 = a1 - b2
	fmt.Println(a1,b2)

	//多重赋值

	q1,w1,e1 := 2,3,4
	var q2 ,w2 ,e2 int  = 100,200,1
	fmt.Println(q1,w1,e1)
	fmt.Println(q2,w2,e2)

	q2,w2,r2 := 11,23,"asd" //11 23 asd  := 多重赋值时如果有新定义的变量，是可以使用自动推导类型:=
	fmt.Println(q2,w2,r2)

	//q2,w2 := 190,21903  没有新定义的变量所以不能使用
	//fmt.Println(q2,w2)
}

func main() {
	var a ,b int = 10,20
	//匿名变量，不接收数据
	_,_,c,d := 100,101,"cccc","dddd"
	fmt.Println(a,b,c,d,)

	//交换变量值
	a,b = b,a
	fmt.Println(a,b)
}
