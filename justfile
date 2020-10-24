test-all:
  go test ./...
lint:
  golangci-lint run
