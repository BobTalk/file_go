package main

import (
	"log"
	"net/http"

	"github.com/autech/file/readpkg"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/read/{type}", readpkg.ReadFileByAddress)

	// r.PathPrefix("/upload/")
	// 指定静态文件目录
	// fs := http.FileServer(http.Dir("./static"))
	// 将静态文件处理器注册到路由
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("服务器启动，访问 http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
