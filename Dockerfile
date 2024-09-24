FROM golang:1.23-alpine AS builder

WORKDIR /app/

COPY . .

WORKDIR /app/

RUN go mod download

RUN go build -o server ./cmd/genproto

FROM builder AS dev

RUN go install github.com/air-verse/air@latest

ENV GOTRACEBACK=all

CMD ["/go/bin/air"]

FROM alpine:3.20

ARG APP_NAME

COPY --from=builder app/server app/server

CMD ["/app/server"]
