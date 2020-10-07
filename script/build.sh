#!/bin/bash

build=$(date +%FT%T%z)
version="v0.0.1"

ldflags="-s -w -X actionflow/config.Build=$build -X actionflow/config.Version=$version"
target="actionflow"

go env -w GOPROXY=https://goproxy.cn,direct

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "$ldflags" -o bin/$target main.go
CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags "$ldflags" -o bin/$target.exe main.go

upx bin/$target
upx bin/$target.exe
