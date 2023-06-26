generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/*.proto

run-server:
	go run ./server/main.go

run-client:
	go run ./client/main.go --name=$(name)