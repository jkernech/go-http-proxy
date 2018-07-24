
build:
	@dep ensure
	@go build ./...

run:
	@go run main.go

release:
	@goreleaser --rm-dist

test:
	@dep ensure
	@go test -cover ./...
