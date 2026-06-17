# Gunakan versi Go yang sama seperti service lainnya
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy file dependency jika ada (go.mod & go.sum)
COPY go.mod ./
# Kalau temanmu gak pakai go.sum atau vendor, baris di bawah ini bisa dimatikan. 
# Tapi kalau ada go.sum, aktifkan lagi:
# COPY go.sum ./ 

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Gunakan image alpine kecil untuk menjalankan binary Go
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]