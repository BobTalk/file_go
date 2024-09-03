package getAbi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/autech/web3Chain/Logs"
	"github.com/gin-gonic/gin"
)

func AbiByType(c *gin.Context) {
	Logs.Println("开始调用接口 getAbi")
	coinType := strings.ToLower(c.Param("type"))
	Logs.Println("开始调用接口 getAbi参数: ", coinType)
	switch coinType {
	case "ozc":
		ReadAbiByFileUri("ozcoinExpandAbi.json", c)
		break
	case "busd":
		ReadAbiByFileUri("ozcoinExpandAbi.json", c)
		break
	case "multiple":
		ReadAbiByFileUri("multiSigWalletAbi.json", c)
		break
	case "toto":
		ReadAbiByFileUri("totoExpandAbi.json", c)
		break
	case "stake":
		ReadAbiByFileUri("ozcoinStakeExpandAbi.json", c)
		break
	case "erc":
		ReadAbiByFileUri("erc20Abi.json", c)
		break
	case "contract":
		ReadAbiByFileUri("contractAbi.json", c)
		break
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "type is defined!"})
		break
	}
}

func ReadAbiByFileUri(fileName string, c *gin.Context) {
	Logs.Println("文件地址 %s : ", fileName)
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 构建文件路径
	absoluteFilePath := filepath.Join(wd, "static", fileName)
	// 读取 JSON 文件
	file, err := os.Open(absoluteFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 将 JSON 内容解析到一个切片中
	var jsonArray []map[string]interface{}
	if err := json.Unmarshal(byteValue, &jsonArray); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 以 JSON 格式返回给前端
	c.JSON(http.StatusOK, gin.H{"data": jsonArray})
}
