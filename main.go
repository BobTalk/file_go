package main

import (
	"log"

	"github.com/autech/web3Chain/getAddress"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getAddress/:type", getAddress.AddressByType)
	log.Println("服务器启动，访问 http://localhost:9100")
	r.Run(":9100")
}
