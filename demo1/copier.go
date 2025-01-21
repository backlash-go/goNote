package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type person struct {
	Name string
	Age  int64
	Sex  int64
}

type child struct {
	Name  string
	Gender1 string


}

func main() {

	p1 := &person{
		Name: "xxb",
		Age:  18,
		Sex:  1,
	}

	var ch child
	err := copier.Copy(&ch, &p1)

	if err != nil{
		fmt.Printf("err is %s\n", err.Error())
	}
	fmt.Printf("res is: %+v\n",ch)

	fmt.Println(ch)
}
