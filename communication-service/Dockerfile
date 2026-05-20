# Stage 1: Builder
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy mod dan sum
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -o communication-service .

# Stage 2: Final Image (Sangat ringan)
FROM alpine:latest
WORKDIR /app

# Copy binary dari stage builder
COPY --from=builder /app/communication-service .

# Expose port (misal 8005)
EXPOSE 8005

CMD ["./communication-service"]