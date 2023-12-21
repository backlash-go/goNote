package main

import (
	"fmt"
	"strings"
)

type Test struct {
	name string
	age  int64
}

func main() {

	//  t1 := Test{
	// 	name: []string{"aa","bb"},
	// 	age: 20000,
	// }
	//
	// fmt.Println(t1)
	//
	//
	//
	//f2(&t1)
	//
	//var x interface{}
	//x = "10"
	//value, ok := x.(string)
	//
	//fmt.Println(value, ",", ok)
	//
	//fmt.Printf("%T\n", x)
	//
	//
	//
	//dict := map[string]int{"key1": 1, "key2": 2,"key3": 3}
	//
	//for k,v := range dict{
	//	fmt.Println(k)
	//	fmt.Println(v)
	//
	//}


   //var sl1 = []string{"admin","devops"}

   sl2 := make([]string,0,0)

   str1 := strings.Join(sl2,",")
   fmt.Println(str1)

	if str1 == "" {
		fmt.Println("aaa")

	}


	fmt.Println(strings.Split(str1,",")[0])


}



func f1() {
	defer fmt.Println("222")
	fmt.Println(1111)
	defer fmt.Println("333")
}
