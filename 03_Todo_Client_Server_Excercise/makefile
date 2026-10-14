generate:
	rm -rf proto_gen && \
	mkdir -p proto_gen && \
	protoc -I=./proto \
	--go_out=./proto_gen \
	--go_opt=paths=source_relative \
	--go-grpc_out=./proto_gen \
	--go-grpc_opt=paths=source_relative \
	./proto/todo.proto && \
	go mod tidy

run-client:
	go run cmd/client/main.go

run-server:
	go run cmd/server/main.go
