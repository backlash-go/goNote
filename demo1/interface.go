package main

import (
	"fmt"
)

/*

接口是一种数据类型，可以接收满足对象的信息
接口是虚的，方法是实在的
接口定义规则， 方法实现规则
接口定义的规则，在方法中必须有对应的实现


一个类型拥有接口的所有方法，那么这个类型实现了这个接口
 */

type HuMan interface {
	SayHello()
}


type student struct {
	name string
	age int
}


type teacher struct {
	name string
	age int
	subject string
}

func (s *student) SayHello(){
	fmt.Println("I am a good student")
	//s.name = "asd"
	fmt.Println(s)
	s.name = "asd"
	fmt.Println(s)
}

func (t *teacher) SayHello(){
	fmt.Println("I am a good teacher")
}

func main() {
   var a HuMan
   stu := student{"xxb",18}
   //fmt.Println(stu)
   a = &stu
   //
   a.SayHello()
   fmt.Printf("%T\n",a)
   //fmt.Println(stu)
   stu.SayHello()
   fmt.Println(stu)



   //tea := teacher{"lyq",10,"math"}



}