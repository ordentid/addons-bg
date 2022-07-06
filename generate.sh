#!/bin/bash

swagger="./www/swagger.json"

if [ -f "$swagger" ] ; then
    rm "$swagger"
fi

# go get -d google.golang.org/protobuf/cmd/protoc-gen-go
# go get -d google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go get -d github.com/mwitkow/go-proto-validators
# go get -d github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
# go get -d github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

# go mod vendor

# go install google.golang.org/protobuf/cmd/protoc-gen-go
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go install github.com/mwitkow/go-proto-validators
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

protoc --proto_path=./proto ./proto/*.proto \
    --proto_path=./proto/libs \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go \
    --plugin=$(go env GOPATH)/bin/protoc-gen-govalidators \
    --go_out=./server/pb --go_opt paths=source_relative \
    --govalidators_out=./server

protoc --proto_path=./proto ./proto/bg_api.proto \
    --proto_path=./proto/libs \
    --proto_path=./vendor \
    --plugin=$(go env GOPATH)/bin/protoc-gen-grpc-gateway \
    --plugin=$(go env GOPATH)/bin/protoc-gen-openapiv2 \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc \
    --go-grpc_out=./server/pb --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./server/pb \
    --grpc-gateway_opt allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
    --openapiv2_out ./proto \
    --openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv

mv ./proto/bg_api.swagger.json ./www/swagger.json

protoc --proto_path=./proto ./proto/bg_gorm_db.proto \
   --proto_path=./proto/libs \
   --plugin=$(go env GOPATH)/bin/protoc-gen-gorm \
   --gorm_out=./server
