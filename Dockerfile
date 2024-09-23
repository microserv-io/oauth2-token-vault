FROM golang:1.23-alpine AS builder

ARG APP_NAME

WORKDIR /app/services

COPY ../core-system/services/oauth .

WORKDIR /app/services/$APP_NAME

RUN go mod download

RUN go build -o server ./cmd/grpc

FROM golang:1.23-alpine AS dev

ARG APP_NAME

WORKDIR /app/services

COPY ../core-system/services/oauth .

WORKDIR /app/services/$APP_NAME

RUN go mod download

RUN go install github.com/air-verse/air@latest

ENV GOTRACEBACK=all

CMD ["/go/bin/air"]

FROM alpine:3.20

ARG APP_NAME

COPY --from=builder app/services/$APP_NAME/server app/server

CMD ["/app/server"]
