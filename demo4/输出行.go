package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/Users/xixianbin/go/src/gobackend/goNote/demo4/TEST")
	if err != nil {
		fmt.Println(err.Error())
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		fmt.Printf("text is :%s\n", input.Text())
	}

}
