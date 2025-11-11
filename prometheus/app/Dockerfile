FROM golang:1.25.3-alpine

WORKDIR /app

COPY main.go go.mod go.sum ./

RUN go mod tidy && \
    go build -o app .

EXPOSE 1234

CMD ["./app"]
