module github.com/gogf/katyusha-demos

go 1.11

require (
	github.com/gogf/gf v1.14.6-0.20201209080429-7316e6648fe3
	github.com/gogf/katyusha v0.0.0-20201209083811-06412389062a
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.29.1
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
)
