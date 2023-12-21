package main

import (
	"fmt"
	"strings"
)

func main() {
sl1 :=[]string{"a","b","c"}
f(sl1...)

s2 := "asdsad"

fmt.Println(strings.Split(s2,s2))

fmt.Println(append([]string{}, "22222"))
fmt.Println(append([]string{}, "3333"))


	
}

func f(s ...string)  {
	fmt.Println(s)
}

func d1(n int) int {
	var sum int
	if n == 1 {
		sum = 1
		return sum
	}
	sum = n + d1(n-1)
	return sum

}


