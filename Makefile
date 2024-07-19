.PHONY: run
run: ## Run server (default: localhost:8080)
	go run ./cmd/apiserver

.PHONY: lint
lint: ## Run golangci-lint (2-minute wait)
	docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v

.PHONY: help
help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
