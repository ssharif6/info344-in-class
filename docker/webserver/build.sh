#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t ssharif6/testserver .
docker push ssharif6/testserver