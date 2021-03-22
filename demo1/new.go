package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var a *int
	if a == nil {
		fmt.Println("a is nil")
	}

	b := new(int)

	if b == nil {
		fmt.Println("b is nil")
	}
	fmt.Println(*b)

	var buf *bytes.Buffer

	buf = new(bytes.Buffer)
	if buf == nil {
		fmt.Println("buf is nil")
	}

	f(buf)

	sl := make([]int,2,2)
	fmt.Println(sl)

	var sl2 []int
	fmt.Println(sl2)
	if sl2 == nil{
		fmt.Println("sl2 is nil")
	}


}


func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}