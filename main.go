package main

import (
	"log"

	"github.com/autech/web3Chain/getAbi"
	"github.com/autech/web3Chain/getAddress"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 信任所有私有网络 IP 段
	// err := r.SetTrustedProxies([]string{"192.168.1.65", "43.207.176.239"})
	// if err != nil {
	// 	log.Println("代理 %s : ", err)
	// }
	r.GET("/getAddress/:type", getAddress.AddressByType)
	r.GET("/getAbi/:type", getAbi.AbiByType)
	log.Println("服务器启动，访问 http://localhost:9100")

	r.Run(":9100")
}
