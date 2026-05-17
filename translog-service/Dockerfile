# Stage 1: Builder
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Stage 2: Final Image
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /main .

# Buka port 8005 khusus untuk Transport/Logistics Service
EXPOSE 8005

CMD ["./main"]