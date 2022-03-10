package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "

	}
	fmt.Println(s)

	var s1 = []string{"sss", "bbb"}

	fmt.Println(strings.Join(s1, "asd"))

	var i1 = []interface{}{"ii1", "ii2"}

	for _, v := range i1 {

		value, ok := v.(int64)

		fmt.Print(value, ok)

		var m1 = map[string]string{"a": "aaa", "b": "bbb", "c": "ccc"}

		for k, v := range m1 {
			fmt.Println(k, v)
		}

		if _, ok := m1["a"]; ok {
			fmt.Println(ok)
		}
	}

	var a int64 = 100
	var b int64 = 100

	if a ==100 && b == 100 {
     fmt.Println("10000")
	}


}
