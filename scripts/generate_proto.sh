#!/bin/bash

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

GRPC_GATEWAY_PATH=`go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2`
GOOGLE_APIS=$GRPC_GATEWAY_PATH/third_party/googleapis

PROTO_FILE_PATHS=`find pkg/proto -name "*.proto"`

for PROTO_FILE_PATH in $PROTO_FILE_PATHS; do

    # Generate grpc stub
    protoc -I /usr/local/include -I . \
        -I $GOPATH/src \
        -I $GOOGLE_APIS \
        --go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out==grpc:. \
        --go-grpc_opt=paths=source_relative \
        $PROTO_FILE_PATH

    # Generate reverse proxy
    protoc -I /usr/local/include -I . \
        -I $GOPATH/src \
        -I $GOOGLE_APIS \
        --grpc-gateway_out . \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        $PROTO_FILE_PATH

    # Generate OpenAPI spec
    protoc -I /usr/local/include -I . \
        -I $GOPATH/src \
        -I $GOOGLE_APIS \
        --openapiv2_out . \
        --openapiv2_opt logtostderr=true \
        $PROTO_FILE_PATH
done
