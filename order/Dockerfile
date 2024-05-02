FROM golang:1.21.9

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /order_service

EXPOSE 8080


CMD ["/order_service"]
