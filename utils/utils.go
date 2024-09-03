package utils

import (
	"os"
	"strconv"

	"github.com/autech/web3Chain/Logs"
	"github.com/joho/godotenv"
)

func isString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}
func GetEnv(key string) string {
	// 加载配置文件
	env := os.Getenv("APP_ENV") // 获取环境变量
	if env == "" {
		env = "development" // 默认环境
	}
	Logs.Println(`当前环境:`, env)
	if err := loadEnv(env); err != nil {
		Logs.Println("Error loading .env file")
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
		Logs.Println("unsupported type: %s", types)
		return nil, nil
	}
}

func loadEnv(env string) any {
	var fileName string
	switch env {
	case "production":
		fileName = ".env"
	case "development":
		fileName = ".env.development"
	default:
		return ""
	}
	err := godotenv.Load(fileName)
	if err != nil {
		Logs.Println("Error loading env file: %v", err)
		return err
	}
	return nil
}
