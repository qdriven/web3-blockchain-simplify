#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
mv main gotest-chain-server
chmod 777 gotest-chain-server
#mv evm-server ~/fsdownload/