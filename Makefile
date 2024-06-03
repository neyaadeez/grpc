clean:
	rm -rf ./proto/dummy_grpc.pb.go
	rm -rf ./proto/dummy.pb.go

goproto: clean
	protoc -Iproto --go_out=. --go_opt=module=github.com/neyaadeez/grpc --go-grpc_out=. --go-grpc_opt=module=github.com/neyaadeez/grpc proto/dummy.proto