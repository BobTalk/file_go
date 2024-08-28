package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	// 加载配置文件
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// 获取并返回
	return os.Getenv(key)
}

// 类型转bool
func TypeTransform(target any, types string) (any, error) {
	switch types {
	case "bool":
		// 断言
		strValue, _ := target.(string)
		return strconv.ParseBool(strValue)
	default:
		fmt.Errorf("unsupported type: %s", types)
		return nil, nil
	}
}
