FROM golang:1.15.13

WORKDIR /app

COPY ./../ .

RUN go mod download

RUN go build ./cmd/server.go