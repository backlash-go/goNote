package main

import (
	"fmt"
	"unsafe"
)

type copyChecker uintptr

func main() {

	var a copyChecker =  100


	b := &a



	fmt.Println(a)

	fmt.Printf("b is %+v\n", uintptr(*b))
	fmt.Printf("unsafeb is %+v\n", uintptr(unsafe.Pointer(b)))



}