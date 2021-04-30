package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	arr1 := []string{"china", "russia", "go"}
	x := MapStrToStr(arr1, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Println(x)
}

func MapStrToStr(arr []string, fn func(s string) string) []string {

	newArray := make([]string, 0)
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i:= 0;i < vdata.Len();i++{
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}



	return result
}
