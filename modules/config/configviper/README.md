# Setup Viper

## Initial setup

```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
go install github.com/go-task/task/v3/cmd/task@v3.25.0
go install gotest.tools/gotestsum@v1.10.0
go install github.com/cosmtrek/air@v1.43.0
go install github.com/spf13/cobra-cli@v1.3.0

cobra-cli init --author "Heitor Carneiro <heitorgcarneiro@gmail.com>" --license MIT --viper

cobra-cli add serve

go run main.go serve
```


## Execution

```
cp .env.example .env

task run
```


## Docs

https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d

https://stackoverflow.com/questions/47185318/multiple-config-files-with-go-viper

