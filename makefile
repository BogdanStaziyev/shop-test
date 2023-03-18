test:
	go test -race -v -cover ./... --count=1

lint:
	golangci-lint run