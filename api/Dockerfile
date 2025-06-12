# Build the builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GO_ENABLED=0 GOOS=linux go build -o /cvmaker-api .

# Build the runtime environment
FROM alpine:latest

WORKDIR /app

COPY --from=builder /cvmaker-api .

EXPOSE 8080

CMD ["/app/cvmaker-api"]
