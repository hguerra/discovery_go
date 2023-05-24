.PHONY: clean test security build run

APP_NAME := main
CURRENT_DIR := $(PWD)
BUILD_DIR := $(CURRENT_DIR)/build
BIN_DIR := $(CURRENT_DIR)/bin

devDependencies:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
	GOBIN=$(BIN_DIR) go install gotest.tools/gotestsum@latest
	GOBIN=$(BIN_DIR) go install github.com/mfridman/tparse@latest
	GOBIN=$(BIN_DIR) go install github.com/DarthSim/hivemind@latest

dependencies:
	go mod download
	go mod tidy

install: devDependencies dependencies

clean:
	rm -rf ./build
	rm -rf ./tmp

lint:
	./bin/golangci-lint run ./...

fmt:
	go fmt ./...

test: clean
	mkdir -p $(BUILD_DIR)/coverage
	APP_PROFILE=test ./bin/gotestsum --junitfile $(BUILD_DIR)/coverage/unit-tests.xml -- -coverprofile=$(BUILD_DIR)/coverage/cover.out -cover ./...
	go tool cover -html=$(BUILD_DIR)/coverage/cover.out -o $(BUILD_DIR)/coverage/coverage.html
	go tool cover -func=$(BUILD_DIR)/coverage/cover.out

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

buildDocker:
	docker build --network host -t $(APP_NAME) .

run:
	go run main.go serve

runDev: clean
	mkdir -p ./tmp/web
	./bin/hivemind Procfile.dev

runDocker:
	docker run -it --rm --network host $(APP_NAME)

BRANCH_TEST := develop01
CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
deployDev01:
		git branch -D $(BRANCH_TEST) || true
		git fetch
		git checkout $(BRANCH_TEST)
		git merge $(CURRENT_BRANCH) --no-edit
		git push origin $(BRANCH_TEST)
		git checkout $(CURRENT_BRANCH)
