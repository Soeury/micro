package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// 服务器端实例

type AddParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type AddResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func Add(x int, y int) int {
	return x + y
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	var param AddParam
	json.Unmarshal(b, &param)
	ret := Add(param.X, param.Y)
	respbytes, err := json.Marshal(AddResult{Code: 200, Data: ret})
	if err != nil {
		log.Fatal(err)
		return
	}

	// 直接返回数据切片
	w.Write(respbytes)
}

func main() {

	http.HandleFunc("/add", addHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
