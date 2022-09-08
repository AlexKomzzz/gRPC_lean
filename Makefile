proto:
	protoc -I=api/proto --go_out=. --go-grpc_out=. api/proto/adder.proto

cleare:
	rm ./pkg/api/adder.pb.go
	rm ./pkg/api/adder_grpc.pb.go

create:
	go build -v ./cmd/server
	go build -v ./cmd/client

start:
	./server

rm:
	rm -f server
	rm -f client