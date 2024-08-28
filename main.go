package main

import (
	"log"
	"net/http"

	"github.com/autech/file/readpkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", Layout)
	r.GET("/read/:type", readpkg.ReadFileByAddress)
	log.Println("服务器启动，访问 http://localhost:3000")
	r.Run(":3000")
}
func Layout(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "首页"})
}
