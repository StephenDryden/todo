.PHONY: build
build:
	GOARCH=amd64 GOOS=linux go build -o bin/bootstrap ./cmd/