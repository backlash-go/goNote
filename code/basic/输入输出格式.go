package main

import "fmt"

func main011() {
	fmt.Println("hello world")
	fmt.Print("aaa\n")
	fmt.Print("bbb\n")

    a := 10
    b := 3.14

    c := 3.14159

    d := "abc"
    //%f默认保存六位小数      默认浮点型数据不是精准的 六位是有效的
	fmt.Printf("%d, %f,%T,%T\n",a,b,b,a) //10, 3.140000,float64,int
    fmt.Printf("%.2f\n",c)     //3.14
    fmt.Printf("%.3f\n",c)     //3.142  会对第四位小数进行四舍五入
    fmt.Printf("%.6f\n",c)     //3.141590
    fmt.Printf("%s\n",d)
}

func main012() {
	var a int = 100
	var b string = "abc"
	s,_ := fmt.Println(b) //0xc00009c008
	fmt.Println(s)  // fmt返回的是字节数，

	fmt.Scan(&a)
	fmt.Println(a)

    //空格 或者 回车表示一个接收的结束

	var c string
	fmt.Scanf("学生姓名：%s",&c)
	fmt.Println(c)

	const a1  = 10
	const a2  = "xxb is a good man"


	fmt.Printf("%d,%T\n",a1,a1) 	//10,int
	fmt.Printf("%s,%T\n",a2,a2) 	//xxb is a good man,string
}

func main() {
	a := 18
	fmt.Printf("%b\n",a)
	fmt.Printf("%o\n",a)
	// x a ~ f      X A ~ F
	fmt.Printf("%x\n",a)
	fmt.Printf("%X\n",a)

    //二进制不能直接表示
	var a1 int = 10
	var a2 int = 010 //8 进制
	var a3 int = 0x10 //16进制
	fmt.Println(a1,a2,a3)
	fmt.Printf("%p",&a1) //内存地址 0xc00009a030

}
