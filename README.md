
# go-grpc-healthcheck
Example project to understand gRPC health-check in Golang

## Services

- [gRPC server](server)
- [gRPC client](client)

---

## Build

1. Server
	```shell
    make build-server
    # or
	go build -o grpc-server ./server
	```

2. Client
	```shell
    make build-client
    # or
	go build -o grpc-client ./client
	```

---

## Run

1. Run server
	```shell
    make run-server
    # or
	./grpc-server
	```

2. In another shell, run client
	```shell
    make run-client
    # or
	./grpc-client
	```

### From source

1. Run server
	```shell
    make run-server-src
    # or
    protoc --proto_path=./proto/ --go_out=plugins=grpc:domain ./proto/*
	GO111MODULE=on go run ./server/main.go
	```

2. In another shell, run client
	```shell
    make run-client-src
    # or
    protoc --proto_path=./proto/ --go_out=plugins=grpc:domain ./proto/*
	GO111MODULE=on go run ./client/main.go
	```

---

## Links

- https://github.com/grpc/grpc/blob/master/doc/health-checking.md
- https://github.com/grpc-ecosystem/grpc-health-probe
- https://godoc.org/google.golang.org/grpc/health
- https://godoc.org/google.golang.org/grpc/health/grpc_health_v1

### tutorials
- https://towardsdatascience.com/grpc-for-production-go-2f62f334824
- https://medium.com/google-cloud/a-simple-http-proxy-for-grpc-healthchecks-be24f78bda7b

### examples
- https://github.com/apssouza22/grpc-server-go
- https://github.com/kelseyhightower/grpc-hello-service
