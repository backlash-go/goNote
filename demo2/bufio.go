package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	fmt.Println(reader.Size())
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line) // the line:http://studygolang.com.
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n)) //It is the home of gophers


	reader2 := bufio.NewScanner(strings.NewReader("abcd\nef\ng"))
	fmt.Println(reader2.Scan(),reader2.Text())
	fmt.Println(reader2.Scan(),reader2.Text())

}
