package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	//p,_ := os.Create("oscreate")
	//p.Write([]byte{97,98,99,100})
	//p.Close()
	p,_ := os.Open("oscreate")
	defer p.Close()
	buf1 := make([]byte,9)
	p.Read(buf1)
	fmt.Println(string(buf1))


file3 := os.NewFile(uintptr(syscall.Stderr),"/dev/stderr")
if file3 != nil{
	file3.WriteString("hello this is a file3 stderr\n")
}


}
