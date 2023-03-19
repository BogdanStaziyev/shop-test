test:
	go test -race -v -cover ./... --count=1

benchmark:
	go test ./test_task2_optimize -bench=. -count=3 -benchmem

lint:
	golangci-lint run