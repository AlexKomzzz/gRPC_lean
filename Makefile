proto:
	protoc -I=api/proto --go_out=. --go-grpc_out=. api/proto/adder.proto
	protoc -I . --grpc-gateway_out ./api/proto \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    api/proto/adder.proto

cleare:
	rm ./pkg/api/*.go

create:
	go build -v ./cmd/server
	go build -v ./cmd/client

start:
	./server

rm:
	rm -f server
	rm -f client