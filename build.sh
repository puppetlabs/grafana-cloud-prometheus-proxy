#!/bin/bash

set -e

echo "Running Go Build"
go get
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web main.go

TIMESTAMP=`date +%s`
docker build --no-cache . -f Dockerfile -t $1:$2

rm web
