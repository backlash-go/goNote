package main

/*
// 发送form-data POST请求
func PostWithFormData(method, url string, postData *map[string]string) map[string]interface{} {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()

	// 构造Post请求
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{Timeout: 2 * time.Second} // 超时2秒
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	fmt.Printf("返回数据%s", data)

	var tempMap map[string]interface{}
	err = json.Unmarshal([]byte(data), &tempMap)
	if err != nil {
		panic(err)
	}

	return tempMap
}
 */
