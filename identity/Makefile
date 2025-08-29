build:
	@go build -o bin/identity ./cmd/api/main.go

run: build
	@./bin/identity

test:
	@go test -v ./...
