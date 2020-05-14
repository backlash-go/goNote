package main

import (
	"fmt"
	"log"
	"reflect"
)
/*

array map slice pointer  .Name() 都是为空
 */
func main() {
	t := reflect.TypeOf("a")
	fmt.Println(t)  //int
	fmt.Println(t.String()) //int
	fmt.Printf("%T\n",t) //*reflect.rtype
	log.Println(t.Kind(),"+++",t.Name(),"---------------")



	v := reflect.ValueOf(100)
	fmt.Println(v) //5
	fmt.Printf("%T\n",v) //reflect.Value

	fmt.Println("+++++")
	fmt.Println(v.Kind())

	if v.Kind() == reflect.Int {
		fmt.Println("yes")
	}
	fmt.Println("+++++")

	vv :=  v.Interface()
	fmt.Printf("%T\n",vv) //int
	fmt.Println(v.String()) //<int Value


	 var sli1  []int
	 var sli2  []string

	 ss1 := reflect.TypeOf(sli1)
	 ss2 := reflect.TypeOf(sli2)

	 //[]int slice
	//[]string slice
	 fmt.Println(ss1,ss1.Kind())
	 fmt.Println(ss2,ss2.Kind())




}
