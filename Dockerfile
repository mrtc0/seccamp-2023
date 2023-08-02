# syntax=docker/dockerfile:1
FROM golang:1.20.6-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

CMD ["./server"]
