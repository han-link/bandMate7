FROM golang:1.26.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./cmd/api

FROM alpine:latest
COPY --from=builder /app/bin /app
ENTRYPOINT ["/app/main"]
