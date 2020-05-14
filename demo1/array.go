package main

import (
	"fmt"
)

func main0003() {
	//array define  默认值为数据类型的默认值
	var a [3]int
	a[0] = 1
	fmt.Printf("%d,%T\n", a[2], a) //0,[3]int

	fmt.Println(len(a), cap(a))
	fmt.Println("+++++")

	var b = [3]int{1, 2, 3}
	fmt.Printf("%d,%T\n", b, b) //[1 2 3],[3]int

	fmt.Println(len(b), len(a))

	for _, v := range b {
		fmt.Println(v)
	}

	e := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(e), e) //5 [1 2 3 4 5]

	//指定一个索引和对应值列表的方式初始化

	type Currency int
	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(USD, symbol[USD])
	fmt.Println(RMB, symbol[RMB])

	f := [...]int{2: 100}
	fmt.Println(f[0], f[1], f[2]) //0 0 100

	//只有当两个数组的所有元素都是相等的时候数组才是相等的
	g := [...]int{1, 2, 3}
	fmt.Println(g == b) //true
	fmt.Println(g == a) //false

	var o *[3]int
	o = &[3]int{1, 2, 3}
	o[1] = 100
	fmt.Println(o[1], o)

	//遍历一维数组
	for i := 0; i < len(o); i++ {
		fmt.Println(o[i])
	}

	//遍历二维数组

	p := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//for i:=0;i<len()
	fmt.Println(len(p), len(p[1]), len(p[2]))
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			fmt.Println(p[i][j])
		}
	}

	//数组的传递是值传递，即数组的副本
	q := p
	fmt.Printf("%v,%T,%p\n", p, p, &p) // [[1 2 3 4] [5 6 7 8] [9 10 11 12]],[3][4]int, 0xc000072060
	fmt.Printf("%v,%T,%p\n", q, q, &q) // [[1 2 3 4] [5 6 7 8] [9 10 11 12]],[3][4]int, 0xc0000720c0

}
