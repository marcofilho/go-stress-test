FROM golang:1.24 as builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]