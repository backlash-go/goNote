package main

import (
	"fmt"
	"sync"
)

func main() {
	var a = sync.Pool{
		New: func() interface{} {return new(a)},
	}

	fmt.Printf()
}
