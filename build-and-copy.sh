#!/bin/bash

# 编译 Go 应用程序
go build -o tmp/web3ChainAddress

# 检查编译是否成功
if [ $? -eq 0 ]; then
  # 创建 tmp 目录（如果不存在）
  mkdir -p tmp
  
  # 拷贝 .env 文件到 tmp 目录
  cp .env tmp/.env
  
  # 拷贝tmp文件到服务器
  scp -r ./tmp root@43.207.176.239:/usr/local/app/tmp
  
  echo ".env 文件已拷贝到 tmp/ 和服务器的 /usr/local/app/tmp/"
else
  echo "编译失败，未拷贝 .env 文件"
  exit 1
fi
