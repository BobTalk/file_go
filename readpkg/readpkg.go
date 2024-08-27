package readpkg

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ReadFileByAddress(c *gin.Context) {
	fileType := c.Param("type")
	switch fileType {
	case "white-paper":
		ReadPDFHandler()
		break
	case "clause":
		GetClause()
		break
	case "privacy-policy":
		GetPrivacyPolicy()
		break
	default:
		break
	}
}

func GetClause() {
	fmt.Sprintf("Clause content here")
}

func GetPrivacyPolicy() {
	fmt.Sprintf("Oz-Privacy-Policy content here")
}

func ReadPDFHandler() {
	// PDF 文件路径
	filePath := "static/test.pdf"

	// 调用 pdfcpu 命令行工具来提取文本
	cmd := exec.Command("pdfcpu", "extract", "text", filePath, "-")
	output, err := cmd.CombinedOutput() // 获取标准输出和错误输出
	if err != nil {
		fmt.Sprintf("Error extracting text from PDF: %s\nOutput: %s", err.Error(), string(output))
		return
	}
}
