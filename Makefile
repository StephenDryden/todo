.PHONY: build
build:
	cd src && GOARCH=amd64 GOOS=linux go build -o ./bin/bootstrap ./cmd/main.go