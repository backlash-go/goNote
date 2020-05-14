package main

import "fmt"

/*
向函数传递参数，传递的是一个副本，要在函数中修改结构体对象的值，需要使用指针
 */

type stu struct {
	name string
	age int
}

func main() {
	s1 := stu{"xxb",18}
    //s1.stu1()
	//fmt.Println(s1)

	s1.stu2()
	fmt.Println(s1)
}

func (s *stu)  stu1(){
     s.name = "lyqxxb"
     fmt.Println(s)
}

func (s stu) stu2()  {
	s.name = "xxb2222"
	fmt.Println(s)
}



func main11() {
	type Point struct{
		X, Y int
	}

	var a *Point
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	if a == nil {
		fmt.Println("a is nil")
	}



	var u1 = &User{"xxb",17}

    u2 := MUser(u1)
    fmt.Println(u1)
    fmt.Println(u2)

    fmt.Printf("%p\n",u1)
    fmt.Printf("%p\n",u2)

    pp := &Point{1,2}
    fmt.Printf("%T\n",pp)

    pp1 := new(Point)
    //pp1 = &Point{4,5}
    *pp1 = Point{4,5}
	fmt.Printf("%T\n",pp1)





}


type User struct {
	name string
	age  int
}

func MUser(u *User )(*User)  {
   u.age = 80
	return u
}
