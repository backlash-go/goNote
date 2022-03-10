package main

import (
	"fmt"
)

type Test struct {
	name []string
	age  int64
}

func main() {

	t1 := &Test{name: []string{"a", "b"}, age: 19}
	fmt.Println(t1)

	var s1 string = "sad"

	var s2 string = fmt.Sprintf("wikaaa%saajek", s1)

	fmt.Println(s2)
	f2(t1)

}

func f2(para *Test) {
	fmt.Println(para.name)

}

func f1() {
	defer fmt.Println("222")
	fmt.Println(1111)
	defer fmt.Println("333")

}
