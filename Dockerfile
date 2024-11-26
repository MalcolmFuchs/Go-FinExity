FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o weather-app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/weather-app .

EXPOSE 8080

CMD ["./weather-app"]