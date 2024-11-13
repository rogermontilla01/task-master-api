FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/*

FROM alpine:latest

COPY --from=builder /app/main /app/

EXPOSE 8080

CMD ["./app/main"]