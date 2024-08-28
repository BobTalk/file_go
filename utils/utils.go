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
	env := os.Getenv("APP_ENV") // 获取环境变量
	if env == "" {
		env = "development" // 默认环境
	}

	if err := loadEnv(env); err != nil {
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

func loadEnv(env string) error {
	var fileName string
	switch env {
	case "production":
		fileName = ".env.production"
	case "development":
		fileName = ".env.development"
	default:
		return fmt.Errorf("unknown environment: %s", env)
	}

	err := godotenv.Load(fileName)
	if err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	return nil
}
