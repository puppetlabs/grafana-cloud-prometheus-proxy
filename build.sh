#!/bin/bash

set -e

echo "Running Go Build"
go get; cd cmd/web; go get; cd ../..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web cmd/web/main.go

TIMESTAMP=`date +%s`
docker build --no-cache . -f Dockerfile -t $1:$2

rm web
