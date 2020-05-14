package main

import "fmt"
/*
接口类型不能直接转换，需要使用断言  也不能参与运算
 */
func test(){
	fmt.Println("test function")
}

func main() {
	var s []interface{}
	s = append(s,123,"sad","sad",test)
	fmt.Printf("%T,%v\n",s,s)
}


func main00() {
	var i interface{}
	i = 10
	fmt.Printf("%T\n",i) //int

	var a int = 10
	fmt.Println(a)
	//fmt.Println(i + a) //invalid operation: i + a (mismatched types interface {} and int)
	//fmt.Println(int(i) + a) //cannot convert i (type interface {}) to type int: need type assertion
}
