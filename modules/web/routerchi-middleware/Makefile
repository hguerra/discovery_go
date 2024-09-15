.PHONY: all install fmt lint clean test build run

install:
	@go install mvdan.cc/gofumpt@v0.7.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	@go install gotest.tools/gotestsum@v1.12.0
	@asdf reshim golang

fmt:
	@gofumpt -l -w .

lint:
	@golangci-lint run ./...

clean:
	@rm -rf ./tmp

test: clean
	@mkdir -p ./tmp/coverage
	@gotestsum --junitfile ./tmp/coverage/unit-tests.xml -- -coverprofile=./tmp/coverage/cover.out -cover ./...
	@go tool cover -html=./tmp/coverage/cover.out -o ./tmp/coverage/cover.html

build:
	@go build -o ./tmp/main ./cmd/api/main.go

run:
	@go run cmd/api/main.go

all: build
