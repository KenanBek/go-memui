.PHONY: test
test:
	go test -v -race ./...

.PHONY: lint
lint:
	golangci-lint run
