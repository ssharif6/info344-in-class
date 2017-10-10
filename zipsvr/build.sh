#!/usr/bin/env bash
echo "building go server for Linux..."
#Linux users, execut: CGO_ENABLED=0 go build -a
GOOS=linux go build 
docker build -t ssharif6/zipsvr .
docker push ssharif6/zipsvr
go clean
