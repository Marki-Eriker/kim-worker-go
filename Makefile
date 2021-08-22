run:
	go run cmd/app/main.go

migrate-up:
	migrate -path internal/database/postgres/migrations -database postgres://postgres:postgres@localhost:5432/kim-worker-go?sslmode=disable up

migrate-down:
	migrate -path internal/database/postgres/migrations -database postgres://postgres:postgres@localhost:5432/kim-worker-go?sslmode=disable down

