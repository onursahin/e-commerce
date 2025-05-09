FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["air"]