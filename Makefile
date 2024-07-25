build:
	@go build -0 bin cmd/main.go

test:
	@go test -v -migrate=true ./...

dev-run:
	@go run cmd/main.go

seed-test-db:
	@go run main.go -seed=true
