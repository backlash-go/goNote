package main

import (
	"fmt"
	"sort"
)

/*
1 使用 var 声明 slice(view slice.go) map 默认都是nil

2 向一个值为nil 的MAP  存入一个元素 会引发panic

3 map 之间不能做比较  唯一的例外是和nil进行比较
 */

func main0() {
	//var m1 map[string]int   m1 为nil

	var m1 = make(map[string]int)  // m1 不为nil
	if m1 == nil{
		fmt.Println("m1 is nil")
	}
	fmt.Println(m1) //map[]

	var m2 = map[string]int{}  //m2 is not nil
	if m2 == nil {
		fmt.Println("m2 is nil")
	}

	m3 := map[int]int{22:20,11:10,5:50,3:30,4:40,66:600}
	fmt.Println(m3,m3[1])

	//delete(m3,3)  //删除某个KEY
	fmt.Println(m3,m3[1])

	//m3[5] = m3[5] + 1 //即使map中不存在5下面的代码也可以正常工作，因为m3[5]失败时将返回0。
    //m3[5] += 1
    //m3[5]++
	fmt.Println(m3)

	for k, v := range m3{
		fmt.Println(k,v)
	}
	/*
	1 10
	2 20
	5 1
	4 40
	*/

	fmt.Println(m3) //map[1:10 2:20 4:40 5:1]

	//判断KEY 是否存在
	if value, ok :=  m3[4]; ok{
		fmt.Println("key is exist")
		fmt.Println(value)

	}else {
		fmt.Println("key is not exist")
		fmt.Println(value)
	}

}

func main1() {
	/*
	1 Map的迭代顺序是不确定的
	 */
	map1 := map[string]int{"xxb":18,"ay":89,"cy":65,"b":99}
	delete(map1,"ay") //map[b:99 cy:65 xxb:18]
	fmt.Println(map1)
	for k, v := range map1{
		fmt.Println(k, v)
	}


	var names []string
	for name := range map1{
		names = append(names,name)
	}
    fmt.Println(names)
	sort.Strings(names)

	fmt.Println(names)

	for _, v := range names{
		fmt.Println(v,map1[v])
	}

    //向一个值为nil 的MAP  存入一个元素 会引发panic
    //var map2 map[int]int
	//map2[1] = 100
	//fmt.Println(map2) // assignment to entry in nil map




}

func main() {
	map1 := map[string]string{"a":"aa","b":"bb","c":"cc"}
	age, ok := map1["aa"]
	fmt.Println(age,ok)
	if age == ""{
		fmt.Println("age is ''")
	}

	//fmt.Println(len(map1),cap(map1))


	seen := make(map[string]bool)
	fmt.Println(seen["iii"])
	fmt.Println(seen["aaa"])

	fmt.Println(!seen["aa"])


	s11 := make([]map[string]string ,0,3)
	fmt.Println(len(s11),cap(s11))

	map2 := map1
	fmt.Printf("%p,%p\n",map2,map1)
	fmt.Println(map2)
	fmt.Println(map1)

	map1["b"] = "bbbbb"

	fmt.Println(map2)
	fmt.Println(map1)

}