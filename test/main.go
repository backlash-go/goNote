package main

import (
	"context"
	"fmt"
	"regexp"
	"time"
)

func modifySlice(s []int) {
	s[0] = 99
	s = append(s, 4)
	fmt.Println("Inside modifySlice:", s)
}

func main() {
	slice := make([]int, 0, 6)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	modifySlice(slice)
	fmt.Println("After modifySlice:", slice)

}

func main1() {
	//git@codeup.aliyun.com:6203423274c3f3b550742191/znzt/xmkf-operation-web.git
	var s string = "git@codeup.aliyun.com:6203423274c3f3b550742191/znzt-sdfaf-fsa_fsaff/xmkf-web-component-lib-vue2-doc.git"
	reg1, _ := regexp.Compile("^git@codeup.aliyun.com:6203423274c3f3b550742191(/[a-zA-Z-]+)+.git$")
	reg2, _ := regexp.Compile("^git@codeup.aliyun.com:6203423274c3f3b550742191.*.git$")

	fmt.Println(reg1.MatchString(s))
	fmt.Println(reg2.MatchString(s))
}

func create() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 在业务代码中使用上下文
	go performOperation(ctx)
}

func performOperation(ctx context.Context) {
	t, f := ctx.Deadline()

	fmt.Println(t, f)
	time.Sleep(7 * time.Second)
	fmt.Println(t, f)

	fmt.Println("11111111")

}
