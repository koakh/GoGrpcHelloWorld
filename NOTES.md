# NOTES

## Links

- [gRPC in Golang: Build a simple API service using gRPC](https://towardsdatascience.com/grpc-in-golang-bb40396eb8b1)
- [Google Protocol Buffers](https://developers.google.com/protocol-buffers)

## Pre-requisites

- [BoostNotes : Go Lang: Install Protocol Buffers and go-support for protobuf](:note:073adef8-47bb-4893-a990-e3700a1ffd28)

## Some NOTES

In gRPC a **client application can directly call methods on a server application on a different machine as if it was a local object**, making it easier for you to create **distributed applications and services**.

## Init Modules

- [Using Go Modules](https://blog.golang.org/using-go-modules)

```shell
$ go mod init koakh.com/helloGrpc
go: creating new go.mod: module koakh.com/helloGrpc
$ cat go.mod 
module koakh.com/helloGrpc

go 1.13
```

now replace `"github.com/grpc-go-course/hello/hellopb"` with `"koakh.com/helloGrpc/hellopb"` in server and client

## Run server and client

```shell
$ go run helloserver/grpcServer.go
Server is listening on 0.0.0.0:50051 ...

$ $ go run helloclient/grpcClient.go
Hello client ...
Receive response => [Hello Jeremy]
```
