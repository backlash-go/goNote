package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//sl1 := make([]string,1,1)
	var sl1 []string

	fmt.Println(sl1,len(sl1),cap(sl1))
	sl1 = append(sl1,"a")
	fmt.Println(sl1,len(sl1),cap(sl1))

	fmt.Println()

   var a,b int = 10, 10
   if b == a {
   	fmt.Println(b)
   }
   fmt.Println(a)

    //s := as{a,b}
	var  map1  =  map[int]int{1:11,2:22}
	//k := 1

	for kk,vv := range map1{
		fmt.Println(kk,vv)
	}
	if k, ok := map1[3];ok{
		fmt.Println(k)

	}else {
		fmt.Println("no")
	}

    fmt.Fprintln(os.Stdout,"sad")

    as1 := as{}

    as2 := as1

    fmt.Printf("%p\n",&as1)
    fmt.Printf("%p\n",&as2)

   // sync.RWMutex{}

   var aaa string
    if aaa == ""{
    	fmt.Println("yyyyes")
	}

	log.Println("sadsad %d",a)
    fmt.Println(Xxb)

}
type as struct {
	a interface{}
	b interface{}
}

type interA interface {

}