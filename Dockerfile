FROM golang:1.23-alpine AS builder

WORKDIR /app/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/grpc

FROM builder AS dev

RUN go install github.com/air-verse/air@latest

ENV GOTRACEBACK=all

CMD ["/go/bin/air"]

FROM alpine:3.20

COPY --from=builder app/server app/server

CMD ["/app/server"]
