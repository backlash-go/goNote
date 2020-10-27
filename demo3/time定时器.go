package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(3 * time.Second)
	//defer t.Stop()

	//t := time.AfterFunc(3 * time.Second, func() {
	//	fmt.Println("222")
	//})

	for  {
		<-t.C
		fmt.Println("11")
	}



	for {
		;
	}



}
