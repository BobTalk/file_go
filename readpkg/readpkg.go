package readpkg

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func ReadFileByAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileType := params["type"]
	switch fileType {
	case "white-paper":
		ReadPDFHandler(w, r)
	case "clause":
		GetClause(w)
	case "privacy-policy":
		GetPrivacyPolicy(w)
	default:
		http.Error(w, "Invalid file type", http.StatusBadRequest)
	}
}

func GetClause(w http.ResponseWriter) {
	fmt.Fprintln(w, "Clause content here")
}

func GetPrivacyPolicy(w http.ResponseWriter) {
	fmt.Fprintln(w, "Oz-Privacy-Policy content here")
}

func ReadPDFHandler(w http.ResponseWriter, r *http.Request) {
	// PDF 文件路径
	filePath := "static/test.pdf"

	// 调用 pdfcpu 命令行工具来提取文本
	cmd := exec.Command("pdfcpu", "extract", "text", filePath, "-")
	output, err := cmd.CombinedOutput() // 获取标准输出和错误输出
	if err != nil {
		http.Error(w, fmt.Sprintf("Error extracting text from PDF: %s\nOutput: %s", err.Error(), string(output)), http.StatusInternalServerError)
		return
	}

	// 将提取的文本内容返回前端
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}
