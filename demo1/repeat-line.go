package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
counts := make(map[string]int)
files := os.Args[1:]

if len(files) == 0 {
	countLines(os.Stdin,counts)
}else {

	for _, arg := range files{
		f, err := os.Open(arg)
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		countLines(f,counts)
	}

}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	//ioutil.ReadFile()
}



func countLines(f *os.File, count map[string]int){

	input := bufio.NewScanner(f)
	for input.Scan(){
		count[input.Text()]++
	}

}
















func main001() {
	count := make(map[string]int)


	//bufio.NewReader(os.Stdin)
	input := bufio.NewScanner(os.Stdin)    //返回一个Scanner 结构体的指针
	for input.Scan() {
		  //counts[line] = counts[line] + 1
		count[input.Text()]++
	    if input.Text() == "ww"{
	    	break
		}
		//if count["ww"] == 3{
		//	break
		//}

	}
	for line, n := range count{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)

		}
	}
}
