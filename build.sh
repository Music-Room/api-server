#!/bin/bash

dep ensure -v

export CGO_ENABLED=0 
export GOOS=linux

go build -a -installsuffix cgo -o bin/server -v

docker build -t music-room .

docker run --rm -ti -p 8080:8080 -p 9000:9000 music-room

#docker build --no-cache -t 848984447616.dkr.ecr.us-east-1.amazonaws.com/music-room .
#`echo $(aws --profile $1 --region eu-west-1 ecr get-login) | sed -e "s/-e none//"` #magic
#docker push 848984447616.dkr.ecr.us-east-1.amazonaws.com/music-room