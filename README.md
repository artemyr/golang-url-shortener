# URL Shortener

## Локалка

### Запуск

в конрне проекта

`make up` - запуск 

`go get github.com/ilyakaznacheev/cleanenv` - установка пакета

`go test ./...` - тесты

___

## Docker
`cd docker`

запустить контейнер `make up`

### Установка зависимостей

в корне проекта

`docker run --rm -v $(pwd):/app -w /app golang:1.25 go mod tidy`

`docker run --rm -v $(pwd):/app -w /app golang:1.25 go get github.com/ilyakaznacheev/cleanenv`

### генерация моков

`cd docker`

`make generate`

### dev режим

`docker compose up go-dev --build`

`docker compose exec go-dev sh`