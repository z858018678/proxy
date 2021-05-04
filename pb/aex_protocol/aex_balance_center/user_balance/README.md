### 编译注意事项

由于本协议首先被 aex_jupiter 框架使用，受框架限制，在编译协议前，需要安装指定版本的 `protoc-gen-*` 文件。
具体操作如下:

#### protoc-gen-go 版本

```bash
git clone https://github.com/golang/protobuf
cd ./protobuf
git checkout v1.3.2
go install ./protoc-gen-go
```

#### protoc-gen-grpc-gateway
```bash
git clone https://github.com/grpc-ecosystem/grpc-gateway
cd ./grpc-gateway
git checkout v1.14.5
go install ./protoc-gen-grpc-gateway
go install ./protoc-gen-swagger
```
