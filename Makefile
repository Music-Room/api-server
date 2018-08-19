BIN=music-room

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFLAGS=-ldflags '-w -s'
DOCKER=docker
DOCKERBUILD=$(DOCKER) build
DOCKERRUN=$(DOCKER) run
DOCKERPUSH=$(DOCKER) push
AWS=alex

GODEP=dep
NEWDEP=$(GODEP) ensure

AWSECR=848984447616.dkr.ecr.us-east-1.amazonaws.com/$(BIN)

all: go-build docker-build docker-push clean

go-build:
	@echo "Golang build executable..."
	$(NEWDEP)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) -o $(BIN) $(GOFLAGS) main.go

docker-build:
	@echo "Docker build service..."
	$(DOCKERBUILD) --no-cache -t $(BIN) .

docker-push:
	@echo "Push Docker image to AWS ECR..."
	docker tag $(BIN):latest $(AWSECR):latest
	eval $(aws ecr --profile alex get-login --no-include-email --region us-east-1 | sed 's|https://||') && docker push $(AWSECR):latest

run:
	@echo "Docker run service..."
	$(DOCKERRUN) \
	--rm \
	-dt \
	-p 8001:8001 \
	--name=$(BIN) \
	$(BIN)

clean:
	@echo "Clean"
	$(GOCLEAN)
	rm -f $(BINARY)
	#docker rmi $(docker images -q)
	#docker rm $(docker ps -q -f status=exited)