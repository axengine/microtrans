module microtrans

go 1.13

replace (
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/google/pprof v0.0.0-20190515194954-54271f7e092f
	golang.org/x/image v0.0.0-00010101000000-000000000000 => github.com/golang/image v0.0.0-20180708004352-c73c2afc3b81
)

require (
	github.com/axengine/go-saga v0.0.0-20200706024708-52fbd8a58b2b
	github.com/axengine/utils v0.0.0-20191028072719-8037441797c3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.13.1
	github.com/micro/go-plugins v1.3.0
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	go.uber.org/zap v1.13.0 // indirect
	google.golang.org/grpc v1.25.1 // indirect
	xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb
)
