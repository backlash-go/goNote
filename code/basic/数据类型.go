package main

import "fmt"

/*
float64 保留小数点后面6位，精度更加精确   float32  保留6位

uint 默认根据操作系统的位数，64位操作系统 即 uint64  uint 如果出现负数 就会溢出变成最大值

虽然 int 会根据操作系统的位数  但是在GO语言中  属于类型不一致,需要使用类型转换
*/

func main001() {
	//uint 默认根据操作系统的位数，64位操作系统 即 uint64
	var a uint = 10
	//无符号uint 如果出现负数 就会溢出变成最大值
	a = a - 20
	fmt.Println(a) //18446744073709551606
}

func main002() {
	//var a int  = 10
	//var b int64  = 20 // invalid operation: a + b (mismatched types int and int64)
	////fmt.Println(a + b) //虽然 int 会根据操作系统的位数  但是在GO语言中  属于类型不一致
	//fmt.Println(a + b)
}

func main003() {
	// bool类型 默认值为false
	var a bool = true
	fmt.Printf("%t\n", a)

	//浮点型
	var b float32 = 30.123
	fmt.Printf("%f\n", b) //30.122999

	var c float64 = 13.25
	fmt.Printf("%f\n", c * 2.2) //29.150000

	//使用自动类型推导 默认是float64
	d := 1.23434
	fmt.Printf("%f,%T\n", d, d) //1.234340,float64

}

func main() {
	// 字符  是uint8 的别名

	var a byte = 'a'
	fmt.Println(a) //ascii 97
    fmt.Printf("%c,%T\n",a,a) //a   uint8

    var b byte = 's'  //ascii 115
    fmt.Println(b)  //115
    fmt.Println(string(b)) //s
    fmt.Println(a + b)  //212

    var c string = "abc"
    fmt.Println(c[1]) //98

    // +32 转大写字母
    var d = 'A'
    fmt.Println(d + 32) // 97
    fmt.Printf("%c\n", d + 32) //a

    var e byte  = '\n'
    fmt.Println(e) // 10
    fmt.Printf("%T,%c\n",e,e) // 回车


}
