build:
	@go build -0 bin cmd/main.go

test:
	@go test ./...

dev-run:
	@go run cmd/main.go


migrations:
	@dbmate up
