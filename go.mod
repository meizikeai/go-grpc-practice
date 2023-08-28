module go-grpc-practice

go 1.21

replace go/protos v1.0.0 => ./protos/proto-practice

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.7.1
	github.com/sirupsen/logrus v1.9.3
	go/protos v1.0.0
	google.golang.org/grpc v1.57.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
