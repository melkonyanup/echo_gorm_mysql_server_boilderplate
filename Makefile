.SILENT:

build:
	go mod download

start:
	go run cmd/app/main.go

test:
	go test --short ./...

docs:
	swag init -g internal/app/app.go

make db.migrate:
	go run cmd/migrations/db_migration.go

test.integration:
	go test ./tests/