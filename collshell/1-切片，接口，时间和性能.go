package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	//fmt.Println(sepIndex)
	fmt.Printf("%p\n",path)


	//dir1 := path[:sepIndex]
	dir1 := path[1:sepIndex:sepIndex]
	fmt.Println(len(dir1),cap(dir1))

	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB


	s := Square{len:5}
	fmt.Printf("%d\n",s.Sides())
	//var  _ Shape = (*Square)(nil)

	var sa *Square = nil

	var  ss2 = (string)("aasdsas")
	//var sa1 *Square

	fmt.Printf("sa is %T,%v\n",ss2,sa)
	var arr1 [5]int
	arr1[0] = 1
	fmt.Printf("%T,%d\n",arr1,len(arr1))

	p  := new(int)
	fmt.Println(*p)
	var i1 int
	fmt.Println(i1)

}


//接口完整性检查

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s Square) Sides()  int {
	return 4

}

//func (s Square) Area()  int {
//	return 4
//
//}