#!/bin/zsh
BUILD_TIME=$(date '+%Y/%m/%d %H:%M:%S')
BUILD_GO_VERSION=$(go version | awk '{print $3"@"$4}')
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
-ldflags \
"-X 'TransGate/version.BuildTime=${BUILD_TIME}' \
-X TransGate/version.BuildGoVersion=${BUILD_GO_VERSION}" \
-o TransGateLinux main.go