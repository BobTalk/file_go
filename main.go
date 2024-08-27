package main

import (
	"log"

	"github.com/autech/file/readpkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/read/:type", readpkg.ReadFileByAddress)

	// r.PathPrefix("/upload/")
	// 指定静态文件目录
	// fs := http.FileServer(http.Dir("./static"))
	// 将静态文件处理器注册到路由
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("服务器启动，访问 http://localhost:3000")
	// http.ListenAndServe(":3000", r)
	r.run(":3000")
}
