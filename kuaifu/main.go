package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 发送一个 GET 请求

	// 基础 URL
	baseURL := "https://apifoxmock.com/m1/4527486-4175086-default/test/login"

	// 构建带有查询参数的 URL
	params := url.Values{}
	params.Add("name", "xxb")

	// 完整的 URL
	fullURL := baseURL + "?" + params.Encode()
	// 发送 GET 请求
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("请求错误:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体错误:", err)
		return
	}

	// 输出响应状态码和响应体
	fmt.Println("状态码:", resp.StatusCode)
	fmt.Println("响应体:", string(body))
}
