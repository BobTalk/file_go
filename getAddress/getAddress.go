package getAddress

import (
	"net/http"

	"github.com/autech/web3Chain/utils"
	"github.com/gin-gonic/gin"
)

func AddressByType(c *gin.Context) {
	params := c.Param("type")
	switch params {
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
	case "preSale":
		c.JSON(http.StatusOK, gin.H{"data": utils.GetEnv("preSaleAddress")})
		break
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "type is defined!"})
		break
	}
}
