package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	t := now.Add(5 * time.Hour)
	fmt.Println(t)

    fmt.Println(t.Sub(now))
}

func main00001() {
	fmt.Println(time.Now())        //2020-05-12 15:16:12.915983 +0800 CST m=+0.000080587
	fmt.Println(time.Now().Unix()) //返回当前时间距离1970年1月1日有多少秒
	fmt.Println(time.Now().UnixNano())

	fmt.Println(time.Now().Second(), time.Now().Hour(), time.Now().Year())

	var a time.Duration = 500 * time.Millisecond
	fmt.Println(a)

	t := time.Unix(1589275722, 0) //时间戳加1个小时的 时间
	fmt.Println(t)

	//时间间隔  time.Duration
	time.Sleep(5) //  休息5纳秒
	// time.Sleep(5 * time.Second)  休息五秒
	fmt.Println("aaaaaaaaa")

	n := 3 //int类型需要转换成 Duration
	time.Sleep(time.Duration(n) * time.Second)

}
