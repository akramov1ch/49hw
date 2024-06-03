run:
	go run cmd/main.go

build:
	go build -o bin/49.Working_on_two_microservices cmd.main.go

export POSTGRES_DB=postgres://postgres:vakhaboff@localhost:5432/shaxboz?sslmode=disable

migrate-file:
	migrate create -ext sql -dir migrations/ -seq 49hw

migrate-up:
	migrate -path migrations -database $(POSTGRES_DB) up

migrate-down:
	migrate -path migrations -database $(POSTGRES_DB) down

migrate-force:
	migrate -path migrations -database $(POSTGRES_DB) force $(version)

.PHONY: run build migrate-up migrate-down migrate-force
