# Default variable
GOIMPORTS ?= goimports
GOCILINT ?= golangci-lint

default:test

test:
	@go test -timeout=10m `go list ./...`

lint:
	@$(GOCILINT) run --no-config --disable=errcheck ./...

gen-version: # Update version
	cd version/scripts && go run gen.go
