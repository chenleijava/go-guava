#!/usr/bin/env bash

# 版本
version="v0.0.6"

# 名字
darwin=table2struct-darwin."$version".bin
linux=table2struct-linux."$version".bin
win=table2struct-win."$version".exe

# 打包
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "$darwin" cli.go
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$linux" cli.go
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "$win" cli.go
