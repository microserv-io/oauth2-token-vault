FROM golang:1.23-alpine AS prepare

WORKDIR /app/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM prepare AS dev

RUN go install github.com/air-verse/air@latest

ENV GOTRACEBACK=all

CMD ["/go/bin/air"]

FROM prepare AS test

CMD ["go", "test", "./...", "-tags", "integration"]

FROM prepare AS builder

RUN go build -o server ./cmd/grpc

FROM alpine:3.20

COPY --from=builder app/server app/server

CMD ["/app/server"]
