package Logs

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Println(message ...any) {
	env := os.Getenv("APP_ENV") // 获取环境变量
	logUrl := "app.log"
	if env != "" {
		logUrl = "/usr/local/app/tmp/app.log"
	}
	// 打开或创建一个日志文件
	file, err := os.OpenFile(logUrl, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("无法打开日志文件: %v", err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)
	// 设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	strSlice := make([]string, len(message))
	for i, name := range message {
		strSlice[i] = fmt.Sprint(name)
	}
	result := strings.Join(strSlice, ": ")
	log.Print(result)
}
