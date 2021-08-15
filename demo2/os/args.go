package main

import (
	"fmt"
	"os"
)

func main() {
	j := 100
	fmt.Println(j)

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("这是第 %d 个参数 %s\n", i, os.Args[i])
	}
}
