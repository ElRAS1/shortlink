all: lint format

lint:
	golangci-lint run 


format:
	find . -name "*.go" -exec go fmt {} \;

.PHONY: lint format
