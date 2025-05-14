DEFAULT_GOAL := run

run:
	@sqlc generate
	@templ generate
	@go fmt
	@go vet
	@go build
	@./aurora
	@go clean

sql:
	sqlc generate