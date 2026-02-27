PROJECT_PATH=$(realpath ../)
UID := $(shell id -u)
GID := $(shell id -g)

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
	docker run --user $(UID):$(GID) --rm -v $(PROJECT_PATH):/app -w /app vektra/mockery