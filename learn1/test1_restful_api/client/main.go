package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 客户端实例

// 不同服务之间的数据传输可以采用接口的形式，需要调用的一方直接调用我们定义好的接口就可以了
//   -1. restful api : 主要用于前后端数据传输，通常采用 http 协议，数据通常以 json 格式数据发送
//   -2. rpc : 主要用于微服务架构下不同的微服务之间进行数据传输，通常采用 tcp 协议，数据以二进制形式传输

// 在微服务的场景下使用 api 的方式来调用远程服务会比较麻烦，所以通常是使用 rpc 的方式

// 那既然有了 restful api , 我们为什么还要使用 rpc 来传输数据呢?
//   因为使用 rpc 能够让我像调用本地函数一样，来调用我们的服务，(os:更方便)

type AddParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type AddResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {

	url := "http://localhost:8080/add"
	param := `{"x":10 , "y":20}`
	resp, err := http.Post(url, "application/json", bytes.NewReader([]byte(param)))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	respbytes, err := io.ReadAll(resp.Body) // 注意服务器端返回的是切片
	if err != nil {
		log.Fatal(err)
		return
	}
	var respData AddResult
	json.Unmarshal(respbytes, &respData)
	fmt.Println(respData.Data) // 30
}
