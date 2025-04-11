# grpc_demo
grpc demo for go

[中文文档](https://zhuanlan.zhihu.com/p/1893702967716213732)

## 1. Install go.mod

```bash 
go get google.golang.org/grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
## 2. Install protobuf

```bash
protoc -I./ --go_out=./src/pb --go-grpc_out=require_unimplemented_servers=false:./src/pb ./src/protofiles/helloworld.proto
```
