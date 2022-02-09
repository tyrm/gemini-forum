fmt:
	go fmt ./...

lint:
	golint ./...

test-local: tidy fmt lint
	go test -cover ./...

test-local-race: tidy fmt lint
	go test -race -cover ./...

test-local-verbose: tidy fmt lint
	go test -v -cover ./...

tidy:
	go mod tidy

.PHONY: fmt lint test