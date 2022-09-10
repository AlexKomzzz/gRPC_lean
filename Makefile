proto:
	protoc -I=api/proto --go_out=./pkg  --go-grpc_out=./pkg api/proto/adder.proto
	protoc -I ./api/proto --grpc-gateway_out ./pkg/api \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    api/proto/adder.proto

	
clear:
	rm ./pkg/api/*.go

create:
	go build -v ./cmd/server

gorun:
	go run ./cmd/server

start:
	./server

rm:
	rm -f server
	rm -f client