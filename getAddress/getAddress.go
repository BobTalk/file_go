package getAddress

import (
	"net/http"
	"strings"

	"github.com/autech/web3Chain/utils"
	"github.com/gin-gonic/gin"
)

func AddressByType(c *gin.Context) {
	coinType := strings.ToLower(c.Param("type"))
	switch coinType {
	case "ozc":
		c.JSON(http.StatusOK, gin.H{"data": utils.GetEnv("ozcAddress")})
		break
	case "multiple":
		c.JSON(http.StatusOK, gin.H{"data": utils.GetEnv("multiSignWalletAddress")})
		break
	case "toto":
		c.JSON(http.StatusOK, gin.H{"data": utils.GetEnv("TotoAddress")})
		break
	case "stake":
		c.JSON(http.StatusOK, gin.H{"data": utils.GetEnv("ozcTotoStakeAddress")})
		break
	case "busd":
		c.JSON(http.StatusOK, gin.H{"data": utils.GetEnv("busdAddress")})
		break
	default:
		break
	}
}
