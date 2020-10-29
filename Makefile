# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# This included makefile should define the 'custom' target rule which is called here.
include $(INCLUDE_MAKEFILE)

build: build_apigateway build_service1 build_service2

build_apigateway:
	@echo "Building apigateway binary..."
	$(GOBUILD) services/apigateway/main.go

build_service1:
	@echo "Building service1 binary..."
	$(GOBUILD) services/service1/main.go

build_service2:
	@echo "Building service2 binary..."
	$(GOBUILD) services/service2/main.go

clean:
	@echo "Cleaning..."
	$(GOCLEAN)

protoc:
	./scripts/generate_proto.sh

protoc_install:
	./scripts/install_protobuf.sh

.PHONY: release
release: custom 