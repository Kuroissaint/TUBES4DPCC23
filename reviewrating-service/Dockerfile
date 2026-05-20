# Stage 1: Builder
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o review-rating-service .

# Stage 2: Final Image
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/review-rating-service .

EXPOSE 8008

CMD ["./review-rating-service"]