package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t1 := "asd"

	fmt.Println(reflect.TypeOf(t1))

	var t2 string = "sadasdsa"

	fmt.Println(reflect.TypeOf(t2))

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	var t3 interface{} = []string{"aaa", "bbb", "ccc"}
	fmt.Println(reflect.ValueOf(t3))
	a := reflect.ValueOf(t3)
	fmt.Println(a.Index(2))
	fmt.Println(-4 / 1)

}
