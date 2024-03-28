FROM golang:alpine AS builder

WORKDIR /opt/app

COPY . .

RUN go build -o crypto-knight-api ./cmd/main.go

FROM alpine

WORKDIR /opt/app

COPY --from=builder /opt/app/crypto-knight-api /opt/app/crypto-knight-api

CMD ["./crypto-knight-api"]