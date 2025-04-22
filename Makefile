.PHONY: serve

serve: ## Start http server
	@go run main.go

allocator: ## Start allocator worker
	@go run main.go allocator:start

test-coverage: ## Show test coverage
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
