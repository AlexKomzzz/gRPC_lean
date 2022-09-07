proto:
	protoc -I=api/proto --go_out=. --go-grpc_out=. api/proto/adder.proto

create:
	go build -v ./cmd/server
	go build -v ./cmd/client

start:
	./server

rm:
	rm -f server
	rm -f client