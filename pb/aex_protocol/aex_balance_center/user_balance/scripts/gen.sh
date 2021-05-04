# golang compile
ImportPath=../proto
OutputPathGolang=../golang

protoc -I. \
    -I$ImportPath \
    -I/usr/local/include/google/protobuf \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:$OutputPathGolang \
    --swagger_out=logtostderr=true:$OutputPathGolang \
    --go_out=plugins=grpc:$OutputPathGolang \
    general.proto \
    producer.proto \
    querier.proto \
    operate_status.proto

