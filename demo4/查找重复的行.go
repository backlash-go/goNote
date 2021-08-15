package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	counts := make(map[string]int)
	f, err := os.Open("TEST")
	content, err := ioutil.ReadAll(f)
	fmt.Println(string(content))

	if err !=nil{
		fmt.Println("sssss")
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {

		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}



}
