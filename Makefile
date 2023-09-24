.DEFAULT_GOAL := all

.PHONY: all

all:
	@go build -o etherscan ./cmd/etherscan

test:
	@go test ./...
