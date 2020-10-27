package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 5; i++ {
		go read(i)
	}

	for i := 1; i < 5; i++ {
		go write(i)
	}
}

func read(index int) {
	fmt.Printf("第%d个读GO", index)
}

func write(index int) {
	fmt.Printf("第%d个写GO", index)
}
