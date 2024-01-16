FROM golang:1.14-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ethereal /app

ENTRYPOINT ["/app/ethereal"]
