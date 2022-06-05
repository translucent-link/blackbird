#! /bin/sh
GOOS=linux GOARCH=amd64 go build -o blackbird.linux . 
docker build --platform=linux/amd64 -t translucentlink/blackbird:$1 .
docker push translucentlink/blackbird:$1