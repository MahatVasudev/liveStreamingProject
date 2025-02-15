build:
	@go build -o ./bin/server ./cmd/main.go

run: build
	@./bin/server

migrate-up:
	@go run ./cmd/migrate/main.go up

migrate-down:
	@go run ./cmd/migrate/main.go down

mirgration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
