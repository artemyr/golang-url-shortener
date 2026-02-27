PROJECT_PATH=$(realpath ../)

up:
	docker compose up go-app --build

dev:
	docker compose up go-dev --build

install:
	docker run --user $(id -u):$(id -g) --rm -v $(PROJECT_PATH):/app -w /app golang:1.25 go mod tidy

in:
	docker compose run --user $(id -u):$(id -g) --rm go-app sh

test:
	docker run --user $(id -u):$(id -g) --rm -v $(PROJECT_PATH):/app -w /app golang:1.25 go test ./...

generate:
	docker run --user $(id -u):$(id -g) --rm -v $(PROJECT_PATH):/app -w /app vektra/mockery