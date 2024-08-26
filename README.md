#### 安装go环境 `go version`

#### 初始化项目  go mod init projectName

#### 检查go环境 go env

#### vscode初始化go仓库所需工具

```
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/lint.git
```

#### 手动创建main.go文件

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
a
```

#### 安装重启服务插件 `监听文件夹和文件类型`

```
go install github.com/air-verse/air@latest

```

#### 运行go项目

```
go run main.go 或者 air
```
