package main

import "fmt"

func main() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("%v\n", p)

		}

	}()

	panic(fmt.Errorf("%s\n", "this is a pannic"))

	fmt.Println("main is over")

	//output :  this is a pannic
}
