module github.com/gogf/katyusha-demos

go 1.11

require (
	github.com/gogf/gf v1.15.5
	github.com/gogf/katyusha v0.1.3-0.20210402065308-c211940af3c6
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.33.2
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)
