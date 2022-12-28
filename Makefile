


.PHONY: run
run:
	go run cmd/server.go


.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/ ./cmd/server.go

.PHONY: api
api:
	protoc --proto_path=./api \
	 	   --go_out=plugins=grpc:./api \
		   ./api/*.proto

