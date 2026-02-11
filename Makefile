PROJECT_PATH=$(realpath .)

up:
	CONFIG_PATH=$(PROJECT_PATH)/config/local.yml go run ./cmd/url-shortener

install:
	go mod tidy