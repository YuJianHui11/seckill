FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o seckill ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/seckill .
COPY --from=builder /app/config/config.yaml .

EXPOSE 8080
CMD ["./seckill"] 