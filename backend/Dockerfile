FROM golang:1.21.6 AS builder

WORKDIR /app/build

COPY . .

RUN go build -o api_server .

EXPOSE 8080

CMD ["./api_server"]
