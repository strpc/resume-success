include .env
env_file = .env
export $(shell sed 's/=.*//' $(env_file))

MAIN_GO = cmd/api/main.go
MIGRATIONS_DIR = db/migrations
POSTGRES_SERVICE_NAME = postgres

.PHONY: run build migrate create_migration up down postgres postgres_down
.DEFAULT_GOAL := run


run:
	go run $(MAIN_GO)

build:
	go build -o cmd/api/api $(MAIN_GO)

up:
	docker-compose up -d

down:
	docker-compose down

postgres:
	docker-compose -f docker-compose.yml run --service-ports -d $(POSTGRES_SERVICE_NAME)

postgres_down:
	docker-compose -f docker-compose.yml rm --stop --force $(POSTGRES_SERVICE_NAME)

migrate:
	migrate -path ./$(MIGRATIONS_DIR) \
	-database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' \
	up

create_migration:
	migrate create -ext sql -dir ./$(MIGRATIONS_DIR) -seq $(name)
