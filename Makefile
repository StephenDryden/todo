.PHONY: build
build:
	GOARCH=amd64 GOOS=linux go build -o /src/bin/bootstrap ./src/cmd/main.go