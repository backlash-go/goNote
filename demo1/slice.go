package main

import "fmt"

/*
       1 切片引用了一个底层的数组
	   2 自动扩容 超过容量扩容为原来的两倍
	   3 一旦扩容重新指向一个新的数组
	   4 切片不存储数据，所以修改数据其实就是修改底层的数组的值
       5 slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。
         不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等（[]byte），
         但是对于其他类型的slice，我们必须自己展开每个元素进行比较：
       6 slice唯一合法的比较操作是和nil比较  但是需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断

*/

func main01() {

	var a = [...]string{1:"a",2:"b",3:"c",4:"d",5:"e",6:"f",7:"g"}
	fmt.Printf("%T\n",a) //[8]string


	s1 := a[1:3]
	fmt.Printf("%d,%d,%v,%T\n",len(s1),cap(s1),s1,s1) //2,7,[a b],[]string
     fmt.Printf("%d,%d,%T\n",a[2:5],a[2:5],a[2:5])
	s2 := a[2:4]
	fmt.Printf("%d,%d,%v,%T\n",len(s2),cap(s2),s2,s2) //2,6,[b c],[]string

	//修改重叠的部分 会改变另外一个切片的VALUE
	s1[1] = "bbbbb"
	fmt.Println(s1,s2) // [a bbbbb]  [bbbbb c]


}

func main02() {
    var s1 = []int{1,2,3,4,5}
	fmt.Println(s1)    //[1 2 3 4 5]
	//reverse(s1)
    //fmt.Println(s1)  //[5 4 3 2 1]

	//reverse(s1[:2])
	//fmt.Println(s1)  //[2 1 3 4 5]

	reverse(s1[2:])
    fmt.Println(s1)    //[1 2 5 4 3]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var s1 = []string{"a","b","c"}
	fmt.Printf("%d,%d,%v,%T\n",len(s1),cap(s1),s1,s1) //3,3,[a b c],[]string

	s2 := s1
	fmt.Printf("%p\n",&s1)
	fmt.Printf("%p\n",&s2)


	//slice唯一合法的比较操作是和nil比较  但是需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断

	/*
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	 */
	var s3  []int
	fmt.Println(s3)
	//if len(s3) == 0{
	if s3 == nil{
		fmt.Println("s3 is nil")
	}

	var s4 = make([]string,3,3) //or   s4 := make([]string,3,3)
	fmt.Printf("%d,%d,%T\n",len(s4),len(s4),s4) //3,3,[]string



    //append
    var runes []rune
	fmt.Printf("%d,%d,%p\n",len(runes),cap(runes),runes)
    fmt.Println("-===========-")
	for _, v := range "hello"{
		runes = append(runes,v)
	}

	fmt.Printf("%d,%d,%p\n",len(runes),cap(runes),runes)
	fmt.Printf("%q\n", runes) //['h' 'e' 'l' 'l' 'o']
	runes = append(runes,100)
	fmt.Printf("%d,%d,%p\n",len(runes),cap(runes),runes)

	var s5 = []int{1,2,3,4,5}
	s6 := make([]int,5,5)   //copy 看的目的切片的len大小
	//var s6 = []int{1,1,1,1}
	fmt.Println(s5,s6) //[1 2 3 4 5] [0 0 0 0 0]

	copy(s6,s5)
	fmt.Println(s5,s6) //[1 2 3 4 5] [1 2 3 4 5]

	fmt.Printf("%p\n",s5)
	fmt.Printf("%p\n",s6)








}
