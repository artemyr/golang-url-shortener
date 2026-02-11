PROJECT_PATH=$(realpath ../)

up:
	docker compose up go-app --build

dev:
	docker compose up go-dev --build

install:
	docker run --rm -v $(PROJECT_PATH):/app -w /app golang:1.25 go mod tidy

in:
	docker compose run --rm go-app sh

test:
	docker run --rm -v $(PROJECT_PATH):/app -w /app golang:1.25 go test ./...

generate:
	docker run --rm -v $(PROJECT_PATH):/app -w /app vektra/mockery