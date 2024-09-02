# Makefile

# Docker Compose ファイルのパス
COMPOSE_FILE=docker-compose.yaml

# Build the Docker images
build:
	docker compose -f $(COMPOSE_FILE) build

# Start the Docker containers
up:
	docker compose -f $(COMPOSE_FILE) up -d

# Stop and remove the Docker containers
down:
	docker compose -f $(COMPOSE_FILE) down

# Install dependencies (if any)
install: app_install front_install

app_install:
	cd app/ && \
	go mod tidy && \
	go mod download && \
	go install github.com/google/wire/cmd/wire@latest  && \
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

front_install:
	cd front/ && yarn install

generate: app_generate front_generate

app_generate:
	cd app/internal/di/ && wire gen
	cd app/ && oapi-codegen -config=config.yaml -package oapi -o internal/handler/oapi/api.gen.go ../schema/api.yaml

front_generate:
	cd front/ && yarn codegen:oapi

# Show logs
logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

# Restart the Docker containers
restart: down up

.PHONY: build up down install clean logs restart