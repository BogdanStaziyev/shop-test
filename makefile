test:
	go test -race -v ./... --count=1

lint:
	golangci-lint run