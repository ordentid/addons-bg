DOCKERCMD=docker

DOCKER_CONTAINER_NAME?=addons-bg-service
DOCKER_CONTAINER_IMAGE?=addons-bg-service:latest
DOCKER_BUILD_ARGS?=
DOCKER_DEBIAN_MIRROR?=http://deb.debian.org/debian

BUILD_DATE?=$(shell date -u +'%Y-%m-%dT00:00:00Z')
BUILD_VERSION?=0.1.0

TOPDIR=$(PWD)
BINARY=addons-bg-service

.FORCE:
.PHONY: build
.PHONY: vet
.PHONY: unit-test
.PHONY: generates
.PHONY: depend
.PHONY: docker-build
.PHONY: solr
.PHONY: clean
.PHONY: install
.PHONY: all
.PHONY: .FORCE

guard-%:
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

build:
	@echo "Executing go build"
	go build -v -buildmode=pie -ldflags "-X main.version=$(BUILD_VERSION)" -o app ./server/
	@echo "Binary ready"

vet:
	@echo "Running Go static code analysis with go vet"
	go vet -asmdecl -atomic -bool -buildtags -copylocks -httpresponse -loopclosure -lostcancel -methods -nilfunc -printf -rangeloops -shift -structtag -tests -unreachable -unsafeptr ./...
	@echo "go vet complete"

unit-test:
	@echo "Executing go unit test"
	go test -v -json -count=1 -parallel=4 ./...
	@echo "Unit test done"

generate:
	go generate ./...

run:
	go run ./server/ grpc-gw-server --port1 9124 --port2 3124 --grpc-endpoint :9124

migrate-db:
	go run ./server/ db-migrate

depend:
	@echo "Pulling all Go dependencies"
	go mod download
	go mod verify
	go mod tidy
	@echo "You can now run 'make build' to compile all packages"

docker-build:
	$(DOCKERCMD) build -t $(DOCKER_CONTAINER_IMAGE) --build-arg GOPROXY=$(GOPROXY) --build-arg GOSUMDB=$(GOSUMDB) --build-arg BUILD_VERSION=$(BUILD_VERSION) $(DOCKER_BUILD_ARGS) .

default: depend

all: depend generate build unit-test

install: depend build

clean:
	rm -f $(BINARY)
	rm -f $(BINARY).exe

proto-gen:
	protoc --proto_path=./proto ./proto/*.proto \
		--proto_path=./proto/libs \
		--plugin=$(go env GOPATH)/bin/protoc-gen-go.exe \
		--plugin=$(go env GOPATH)/bin/protoc-gen-govalidators.exe \
		--go_out=./server/pb --go_opt paths=source_relative \
		--govalidators_out=./server
	protoc --proto_path=./proto ./proto/bg_api.proto \
		--proto_path=./proto/libs \
		--plugin=$(go env GOPATH)/bin/protoc-gen-grpc-gateway.exe \
		--plugin=$(go env GOPATH)/bin/protoc-gen-openapiv2.exe \
		--plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc.exe \
		--go-grpc_out=./server/pb --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./server/pb \
		--grpc-gateway_opt allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
		--openapiv2_out ./proto \
		--openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv
	mv ./proto/bg_api.swagger.json ./www/swagger.json
	protoc --proto_path=./proto ./proto/bg_gorm_db.proto \
	--proto_path=./proto/libs \
	--plugin=$(go env GOPATH)/bin/protoc-gen-gorm.exe \
	--gorm_out=./server
