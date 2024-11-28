build:
	@go build -o bin/sso-service ./cmd/server/main.go

run: build
	@./bin/sso-service --config=./config/local.yaml

test:
	@go test -v ./...

migrate:
	@go run ./cmd/migrate/main.go --config=./config/local.yaml --migrations-path=./migrations
