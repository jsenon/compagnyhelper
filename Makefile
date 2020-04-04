#-----------------------------------------------------------------------------
# Global Variables
#-----------------------------------------------------------------------------

DOCKER_USER ?= $(DOCKER_USER)
DOCKER_PASS ?= 

DOCKER_BUILD_ARGS := --build-arg HTTP_PROXY=$(http_proxy) --build-arg HTTPS_PROXY=$(https_proxy)

APP_VERSION := latest
PACKAGE ?= $(shell go list ./... | grep configs)
VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
GIT_COMMIT=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

GOLINTER:=$(shell command -v golangci-lint 2> /dev/null)

override LDFLAGS += \
  -X ${PACKAGE}.Version=${VERSION} \
  -X ${PACKAGE}.BuildDate=${BUILD_DATE} \
  -X ${PACKAGE}.GitCommit=${GIT_COMMIT} \


#-----------------------------------------------------------------------------
# BUILD
#-----------------------------------------------------------------------------

.PHONY: default build test publish build_local lint artifact_linux artifact_darwin deploy
default:  test lint build swagger

test:
	go test -v ./...
build_local:
	go mod tidy
	go build -ldflags '${LDFLAGS}' -o compagnyhelper  
build:
	go mod tidy
	docker build $(DOCKER_BUILD_ARGS) -t $(DOCKER_USER)/compagnyhelper:$(APP_VERSION) -f ./build/Dockerfile .
lint:
ifndef GOLINTER
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.15
endif
	golangci-lint run

swagger:
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
	GO111MODULE=off go generate

artifact_linux:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '${LDFLAGS}' -o compagnyhelper-linux

artifact_darwin:
	CGO_ENABLED=0 GOOS=darwin go build -ldflags '${LDFLAGS}' -o compagnyhelper-darwin

#-----------------------------------------------------------------------------
# DEPLOY
#-----------------------------------------------------------------------------
deploy:
	kubectl apply -f deployment/
#-----------------------------------------------------------------------------
# PUBLISH
#-----------------------------------------------------------------------------

.PHONY: publish 

publish: 
	docker push $(DOCKER_USER)/compagnyhelper:$(APP_VERSION)

#-----------------------------------------------------------------------------
# CLEAN
#-----------------------------------------------------------------------------

.PHONY: clean 

clean:
	rm -rf compagnyhelper
