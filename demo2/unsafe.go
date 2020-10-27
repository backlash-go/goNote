package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

func main() {

	//返回一个类型所占用的内存大小，这个大小只有类型有关，和类型对应的变量存储的内容大小无关，比如bool型占用一个字节、int8也占用一个字节。
	str := "hello"
	fmt.Println(unsafe.Sizeof(str))



	var str1 string
	fmt.Printf("%T\n",str1)

	pint := (*int)(unsafe.Pointer(&str1))
	fmt.Printf("%T\n",pint)

	var ua unsafe.Pointer

	fmt.Printf("%T\n",ua)

    var ad int32 = 10

	//fmt.Println(atomic.AddInt32(&ad,1))


    fmt.Println(atomic.LoadInt32(&ad))

    //var a *int32
    fmt.Println(new(*int32))


	//fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"



}
