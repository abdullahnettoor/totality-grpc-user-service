proto:
	protoc --go_out=. --go-grpc_out=. pb/*.proto

server:
	go run cmd/main.go

new_client:
	go run client/client.go