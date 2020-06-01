package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	src := strings.NewReader("abcdefg")
	dst := new(strings.Builder)

	written,_ := io.CopyN(dst,src,3)
	fmt.Println("src is :  ",  src.Len())

	fmt.Println("dst is : ",  dst.String())
	fmt.Println("dst has been written: ",written)
}
