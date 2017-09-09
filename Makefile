
build:
	@dep ensure
	@go build ./...

run:
	@go run main.go

test:
	@dep ensure
	@go test -cover ./...
