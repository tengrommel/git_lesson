package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello, world!")
}

// 入口函数
func main() {
	http.HandleFunc("/", handler)
	// 监听所有网卡
	http.ListenAndServe(":8080", nil)
}
