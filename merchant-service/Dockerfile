FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

# Diubah ke 8081 sesuai dengan port di main.go merchant-service
EXPOSE 8081

CMD ["./main"]
