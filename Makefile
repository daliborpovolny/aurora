DEFAULT_GOAL := run

run:
	sqlc generate
	templ generate
	go fmt
	go vet
	go build
sql:
	sqlc generate
templ:
	templ generate
air:
	air -c air.toml
clean:
	go clean
