clean:
	rm -rf ./proto/*.pb.go

goproto: clean
	protoc -Iproto --go_out=. --go_opt=module=github.com/neyaadeez/grpc --go-grpc_out=. --go-grpc_opt=module=github.com/neyaadeez/grpc proto/*.proto

cleanbin:
	rm -rf ./bin/*

gobuild: cleanbin
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client