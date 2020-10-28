#!/bin/bash

go get github.com/grpc-ecosystem/grpc-gateway
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

GRPC_GATEWAY_PATH=`go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway`
GOOGLE_APIS=$GRPC_GATEWAY_PATH/third_party/googleapis

PROTO_FILE_PATHS=`find pkg/proto -name "*.proto"`

for PROTO_FILE_PATH in $PROTO_FILE_PATHS; do
    # Generate grpc stub
    protoc -I/usr/local/include -I. \
        -I$GOPATH/src \
        -I$GOOGLE_APIS \
        --go_out=plugins=grpc:. \
        $PROTO_FILE_PATH

    # Generate reverse proxy
    protoc -I/usr/local/include  -I. \
        -I$GOPATH/src \
        -I$GOOGLE_APIS \
        --grpc-gateway_out=logtostderr=true:. \
        $PROTO_FILE_PATH --swagger_out=json_names_for_fields=true:.
done
