module proxy

go 1.15

replace github.com/douyu/jupiter v0.2.5 => ../../aex_jupiter

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/alibaba/sentinel-golang v1.0.2 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/douyu/jupiter v0.2.5
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-resty/resty/v2 v2.4.0 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/automaxprocs v1.3.0 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.26.0
)
