# Stage 1: Builder
FROM golang:1.21-alpine AS builder

# Install git untuk dependency jika diperlukan
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod dan sum dulu agar layer cache efisien
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary (ganti {{SERVICE_NAME}} dengan nama service, misal: order-service)
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Stage 2: Final Image
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /main .

# Expose port (nanti di-override saat running sesuai jatah port)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]