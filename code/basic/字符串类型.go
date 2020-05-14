package main

import "fmt"

/*
go语言中一个汉字 占用三个字节
 */

func main0001() {
	ch := 'a'
	var cha1  byte = 'b'
	var cha2  byte = 19
	str := "a"

	fmt.Printf("%T\n",ch)  //int32
	fmt.Println(ch)
	fmt.Printf("%T\n",str) // == 'a''\0'   string  '\0'表示字符串的结束标志
	fmt.Printf("%T\n",cha1)  //uint8
	fmt.Printf("%T\n",cha2)  //uint8
}

func main() {
	//go语言中一个汉字 占用三个字节
	a := "创指博客"
	fmt.Println(len(a)) //12

	str1 := "你好"
	str2 := "老铁"
	str3 := str1 + str2
	fmt.Println(str3)  //你好老铁

}
