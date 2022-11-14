package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world")
	fmt.Fprintln(w, "http server is running!")
}

func main() {
	fmt.Println("http server")
	// http://localhost:8080/index   这就类似spring的controller层，映射路径和方法的对应关系
	http.HandleFunc("/index", IndexHandler)
	// 在8080端口启动服务
	http.ListenAndServe(":8080", nil)
}
