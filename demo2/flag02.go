package main

import (
	"flag"
	"fmt"
)

func main() {
	//s := flag.String("conf","./conf.yaml","usage")
	//flag.Parse()
	//fmt.Println(*s)

	//var flagvar int
	//flag.IntVar(&flagvar,"flagname",12,"usage is")
	//fmt.Println(flagvar)

	var flagvar string

	flag.StringVar(&flagvar,"flagname","baby","usage is")
	flag.Parse()
	fmt.Println(flagvar)



}
