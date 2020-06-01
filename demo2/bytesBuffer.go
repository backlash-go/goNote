package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer

	contents := "Simple byte buffer for marshaling data."
	contents2 := "ab"
	//1 :
	//buffer1.WriteString(contents)
	//fmt.Println(buffer1.Len(), buffer1.Cap())
	//
	//p1 := make([]byte, 7)
	//n, _ := buffer1.Read(p1)
	//fmt.Println(n)
	//fmt.Println(buffer1.Len(), buffer1.Cap())
	//fmt.Println(p1)
	//fmt.Println(buffer1.String())

	// 2:
	buffer1.WriteString(contents)
	fmt.Println(buffer1.Len(), buffer1.Cap())
	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes were read. (call Read)\n", n) // 0
	fmt.Printf("The length of buffer: %d\n", buffer1.Len()) //0
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap()) //0

	fmt.Println(string(p1))

	//3
	buffer2 := bytes.NewBufferString(contents2)
	unreadBytes := buffer2.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) // 未读内容为：[97 98]。

}
