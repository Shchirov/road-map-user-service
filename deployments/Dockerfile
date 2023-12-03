FROM golang:1.19-alpine

WORKDIR /app

COPY .. .

RUN go mod download

RUN go build -o road-map-user cmd/main.go

EXPOSE 9090

CMD ["./road-map-user-service"]