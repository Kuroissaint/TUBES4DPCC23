# Stage 1: Builder
FROM golang:1.26-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum duluan biar cache Docker optimal
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code dari folder shop-order-service
COPY . .

# Build binary Go-nya
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Stage 2: Final Image (Pakai Alpine murni biar super ringan)
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Ambil hasil build dari stage 1
COPY --from=builder /main .

# Buka port 8009 khusus untuk Shopping Order Service
EXPOSE 8009

# Jalankan aplikasinya
CMD ["./main"]