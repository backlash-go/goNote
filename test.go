package main

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	//go f1()
	//go f2()
	//a := 10

	ss1 := make([]int, 3, 5)
	fmt.Printf("%v,%p,%p\n", ss1, ss1, &ss1)
	ss1 = append(ss1, 1)
	fmt.Printf("%v,%p,%p\n", ss1, ss1, &ss1)

	ss1 = append(ss1, 11)
	fmt.Printf("%v,%p,%p\n", ss1, ss1, &ss1)

	ss1 = append(ss1, 111)
	fmt.Printf("%v,%p,%p\n", ss1, ss1, &ss1)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)

	log.Println("sad")

	//

	var s5 []int
	fmt.Println(unsafe.Sizeof(s5)) //24  字节
	fmt.Println(s5)
	if s5 == nil {
		fmt.Println("s5 is nil")
	}

	var ss5 *[]int
	fmt.Println(unsafe.Sizeof(ss5)) //8

	var a11 int = 10
	b11 := a11
	fmt.Println(a11, b11)

	b11 = 100
	fmt.Println(a11, b11)

	var s string = "111"
	s += "a"
	fmt.Printf("%s,%p\n", s,&s)
	s += "b"
	fmt.Printf("%s,%p\n", s,&s)

	var sss1 []int
	fmt.Println(len(sss1),cap(sss1))

	var aaa1 [3]int
	fmt.Println(len(aaa1))

	ta(1,"s")

}

func  ta (arr  ...interface{}){
	fmt.Println(arr)
	fmt.Printf("%T\n",arr)
}
