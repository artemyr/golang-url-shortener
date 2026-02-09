установить какую-нибудь зависимость
в корне проекта
docker run --rm -v $(pwd):/app -w /app golang:1.25 go mod tidy
docker run --rm -v $(pwd):/app -w /app golang:1.25 go get github.com/ilyakaznacheev/cleanenv

зайти посмотреть в контейнер
docker compose run --rm go-app sh