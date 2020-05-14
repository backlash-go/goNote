package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	period := flag.Duration("period", 1 * time.Second, "period sleep")
	//flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println("hello")
}
