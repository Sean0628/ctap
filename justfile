test:
  go test ./...
fmt:
  go fmt ./...
lint:
  golangci-lint run -E gofmt
