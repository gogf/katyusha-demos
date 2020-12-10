module github.com/gogf/katyusha-demos

go 1.11

require (
	github.com/gogf/gf v1.14.6-0.20201210171253-363f6eba4454
	github.com/gogf/katyusha v0.0.0-20201210171714-50a4d6c5c88f
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.29.1
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
)
