.PHONY: proto build test go-lint proto-lint


# Main go commands

build:
	go build ./...
test:
	go test -cover ./...
go-lint:
	golangci-lint run ./...

# Main protobuf commands

proto:
	buf generate
proto-lint:
	buf lint



