# Build Stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app .

# Runtime Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]