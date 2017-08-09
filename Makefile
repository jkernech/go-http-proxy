
build:
	@go build ./...

run:
	@go run main.go

test:
	@go test -cover ./...
