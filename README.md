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

[参考文档地址](https://www.topgoer.com/)

#### 检查文件类型
```
file /usr/local/app/tmp/web3ChainAddress
```

#### 起服务
```
sudo systemctl daemon-reload
sudo systemctl start yourapp
sudo systemctl enable yourapp
```
#### 添加配置文件app.service
```
/etc/systemd/system/yourapp.service
[Unit]
Description=Your Go App

[Service]
ExecStart=/path/to/yourapp
WorkingDirectory=/path/to/working/directory
Restart=always
User=www-data
# Group=www-data
Environment=PORT=8080

[Install]
WantedBy=multi-user.target

```

#### 错误状态码
```
1. 203  执行文件类型不对
```