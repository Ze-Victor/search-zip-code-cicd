.PHONY: default run build test lint docs clean

APP_NAME=search-zip-code

default: run

run:
	@cd cmd/api && go run routes.go main.go

build:
	@go build -o $(APP_NAME) cmd/api/main.go

test:
	@cd internal/services && go test -v

lint:
	@golangci-lint run

docs:
	@swag init -g cmd/api/main.go

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs