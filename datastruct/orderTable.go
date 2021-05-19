package main

import "fmt"

const maxSize = 5

func main() {
	//els := [maxSize]ElemType{{name:"xxb1",tel:"1"},{name:"xxb2",tel:"2"}}
	//sq1 := SqlList{data:els,length:2}
	sq1 := SqlList{}
	//sq1.ListInsert()
	sq1.ListInsert(1, ElemType{name: "bob1", tel: "0001"})
	sq1.ListInsert(1, ElemType{name: "bob2", tel: "0002"})
	sq1.ListInsert(1,ElemType{name:"bob3",tel:"0003"})
	sq1.ListInsert(1,ElemType{name:"bob4",tel:"0004"})
	sq1.ListInsert(1,ElemType{name:"bob5",tel:"0004"})
	//sq1.ListInsert(1,ElemType{name:"bob4",tel:"0004"})
	fmt.Println(sq1)

	sq1.ListDelete(3)
	fmt.Println(sq1)

}

type ElemType struct {
	name string
	tel  string
}

type SqlList struct {
	data   [maxSize]ElemType
	length int
}

//1 可插入的位置
//2 什么情况下不能再插入数据
//3 每次插入元素时移动元素的次数

func (s *SqlList) ListInsert(i int, el ElemType) bool {
	// 判断顺序表是为满
	if s.length == len(s.data) {
		fmt.Printf("顺序表长度为%d 无法插入\n", s.length)
		return false
	}
	//什么情况下不能再插入数据 因为插入数据需要连续所以 s.length+1 < i
	if s.length+1 < i || i < 1 {
		fmt.Printf("顺序表长度为%d 超出范围\n", s.length)
		return false
	}
	//逻辑序号转为物理序号    1 2 3 4 5 <--> 0 1 2 3 4
	i--
	for j := s.length; j > i; j-- {
		s.data[j] = s.data[j-1]
	}
	s.data[i] = el
	s.length++
	return true
}

func (s *SqlList) ListDelete(i int) bool {
	//判断顺序表是否为空
	if s.length == 0 {
		fmt.Printf("顺序表长度为%d 无法插入\n", s.length)
		return false
	}

	if i > s.length || i < 1 {
		return false
	}
	//逻辑序号转为物理序号    1 2 3 4 5 <--> 0 1 2 3 4


	i--

	for j := i; j < s.length-1; j++ {
		s.data[j] = s.data[j+1]
	}
	s.data[s.length-1] = ElemType{}

	s.length--
	return true
}
