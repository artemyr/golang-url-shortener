# ---------- build stage ----------
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/url-shortener

# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8082

CMD ["./app"]
