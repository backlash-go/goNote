package main

import (
	"fmt"
	"unsafe"
)

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

func main() {
	fmt.Println(unsafe.Sizeof(string("abcdef")))

	var str1 string
	fmt.Printf("%T\n",str1)

	pint := (*int)(unsafe.Pointer(&str1))
	fmt.Printf("%T\n",pint)


	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"



}
