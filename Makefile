.PHONY: run
run:
	go run ./cmd/apiserver

.PHONY: lint
lint:
	docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v
