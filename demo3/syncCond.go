package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"sync"
)

func main() {
	//var mailbox uint8
	//var lock sync.RWMutex
	//sendcond := sync.NewCond(&lock)
	//recv := sync.NewCond(lock.RLocker())

	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	reCond := sync.NewCond(lock.RLocker())

	fmt.Println("ok")

}
