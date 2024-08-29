package getAbi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AbiByType(c *gin.Context) {
	coinType := strings.ToLower(c.Param("type"))
	switch coinType {
	case "ozc":
		ReadAbiByFileUri("static/ozcoinExpandAbi.json", c)
		break
	case "multiple":
		c.JSON(http.StatusOK, gin.H{"data": ""})
		break
	case "toto":
		c.JSON(http.StatusOK, gin.H{"data": ""})
		break
	case "stake":
		c.JSON(http.StatusOK, gin.H{"data": ""})
		break
	case "busd":
		c.JSON(http.StatusOK, gin.H{"data": ""})
		break
	}
}

func ReadAbiByFileUri(uri string, c *gin.Context) {
	// 读取 JSON 文件
	file, err := os.Open(uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

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
	c.JSON(http.StatusOK, jsonArray)
}
