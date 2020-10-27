package main

import "fmt"

type Pet interface {
	show()
}

type Dog struct {
	name string
}

func (a Dog) show() {
	return
}

func main() {
	var dog *Dog
	if dog == nil {
		fmt.Println("dog is nil")

	}
	fmt.Printf("%p,%T,%v\n", dog, dog, dog)

	var a int64
	a = 1
	b := &a
	fmt.Printf("%p,%T\n", &a, &a)
	fmt.Printf("%p,%T\n", b, b)
	fmt.Printf("%p,%T\n", a, a)

	var inter Pet = dog
	if inter == nil {
		fmt.Println("inter is nil")

	}

	fmt.Printf("%p,%T\n", inter, inter)

	m1 := make(map[int]string)
	m1[1] = "a"
	m1[2] = "b"

	if v, ok := m1[3]; !ok {
		fmt.Printf("%T,%v\n", v, v)
		fmt.Println(v, ok)
	}

	var m2  map[int]int
	m2[1] = 2
	fmt.Println(m2)


}
