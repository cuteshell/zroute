proto:
	protoc --proto_path=common/proto --go_out=common/proto/gen --go-grpc_out=common/proto/gen memdb.proto

mmdb:
	go build -mod=vendor -o build/bin/mmdb cmd/core/mmdb/mmdb.go

zroutec:
	go build -mod=vendor -o build/bin/zroutec cmd/zroutec/zroutec.go

all: mmdb zroutec
