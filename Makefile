build:
	@go build -0 bin cmd/main.go

test:
	@go test ./...

dev-run:
	@go run cmd/main.go

seed-test-db:
	@go run main.go -seed=true
