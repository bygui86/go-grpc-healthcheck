
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## code

build-protobuf :		## Compile protobuf
	protoc --proto_path=./proto/ --go_out=plugins=grpc:domain ./proto/*

build-server : build-protobuf		## Build server
	go build -o grpc-server ./server

build-client : build-protobuf		## Build client
	go build -o grpc-client ./client

run-server : build-server		## Run server
	./grpc-server

run-client : build-client		## Run client
	./grpc-client

run-server-src : build-protobuf		## Run server
	GO111MODULE=on go run ./server/main.go

run-client-src : build-protobuf		## Run client
	GO111MODULE=on go run ./client/main.go

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo ""
