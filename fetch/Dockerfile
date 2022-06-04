# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY swaggerui ./swaggerui

EXPOSE 8081
CMD [ "/app/main" ]