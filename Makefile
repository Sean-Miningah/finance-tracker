build:
	@go build -0 bin cmd/main.go

test:
	@go test -v ./...

dev-run:
	@go run cmd/main.go

migrate:
	@dbmate up
