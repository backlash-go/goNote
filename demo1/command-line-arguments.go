package main

import (
	"fmt"
	"os"
	"strings"
)


func main0001() {
	// go run demo1/command-line-arguments.go 1 2 3
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]    //如果连接涉及的数据量很大，这种方式代价高昂   s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)  //1 2 3
}

func main0002() {
// go run demo1/command-line-arguments.go 1 2 3
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
//1
//2
//3
	}

	fmt.Println(os.Args[0]) // /var/folders/_7/xlp5f0rj3f3fk6x56jghchqc0000gn/T/go-build426291999/b001/exe/command-line-arguments



//死循环的两种方式
	//for true {
	//	fmt.Println("a")
	//}

	//for {
	//	fmt.Println("s")
	//}
}

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
}
