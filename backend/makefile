SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=test/code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api
# 	GOARCH=amd64 $(GOCMD) build -o $(BINARY_DIR)/api-linux-amd64 -v ./cmd/api/main.go

build-run: build ## run project build file if not exist build it
#	./$(BINARY_DIR)/api-linux-amd64
	./$(BINARY_DIR)/api

run: ## Start application
	$(GOCMD) run ./cmd/api/main.go

air: ## Start application
	air -c .air.toml

start:
	@make docker-up
	@make air

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out -o $(CODE_COVERAGE).html
	open $(CODE_COVERAGE).html

deps: ## Install dependencies
#	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
#	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
#	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swagger: ## install swagger and its dependencies for generate swagger using swag
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest
	$(GOCMD) get -u github.com/swaggo/swag/cmd/swag
	$(GOCMD) get -u github.com/swaggo/gin-swagger
	$(GOCMD) get -u github.com/swaggo/files

swag: ## Generate swagger docs
	swag init -g pkg/api/server.go -o ./cmd/api/docs

format: ## To check the code standard violations and errors
	golangci-lint run --fix

mockgen: # Generate mock files for the test
	mockgen -source=pkg/repository/interfaces/auth.go -destination=pkg/repository/mock/auth_mock.go
	mockgen -source=pkg/repository/interfaces/order.go -destination=pkg/repository/mock/order_mock.go
	mockgen -source=pkg/repository/interfaces/product.go -destination=pkg/repository/mock/product_mock.go
	mockgen -source=pkg/repository/interfaces/user.go -destination=pkg/repository/mock/user_mock.go
	mockgen -source=pkg/service/token/token.go -destination=pkg/service/mock/token/token_mock.go
	mockgen -source=pkg/service/cloud/cloud.go -destination=pkg/service/mock/cloud/cloud_mock.go
	mockgen -source=pkg/usecase/interfaces/auth.go -destination=pkg/usecase/mock/auth_mock.go
	mockgen -source=pkg/usecase/interfaces/order.go -destination=pkg/usecase/mock/order_mock.go
	mockgen -source=pkg/usecase/interfaces/product.go -destination=pkg/usecase/mock/product_mock.go
	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/usecase/mock/user_mock.go

docker-up: ## To up the docker compose file
	docker compose up -d

docker-down: ## To down the docker compose file
	docker compose down

docker-build: ## To build newdocker file for this project
	docker build -t nikhil382/ecommerce-gin-clean-arch .

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

seed:
	go run cmd/seed/seeder.go
