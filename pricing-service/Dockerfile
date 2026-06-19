FROM golang:1.26-alpine AS builder
WORKDIR /app

# Matikan Go Workspaces biar build lancar
ENV GOWORK=off

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

# Ubah ke 8087 sesuai dengan main.go pricing-service
EXPOSE 8087
CMD ["./main"]