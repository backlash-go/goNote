package main

import "fmt"

func main() {
	var i interface{}
	i = 10
	v, ok := i.(int)
	fmt.Println(v, ok)

	var a int = 100
	fmt.Println(v + a + 10)

	var s []interface{}
	s = append(s, 123, "sad", 44, "sad", test1)

	for _, v1 := range s {
		if data, ok := v1.(int); ok {
			fmt.Println(data)
		} else if data, ok := v1.(func()); ok {
			data()
		}
	}

}
func test1() {
	fmt.Println("test function")
}
