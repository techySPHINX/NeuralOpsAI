FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY cmd/api/*.go ./cmd/api/

RUN go build -o /neuralops-api ./cmd/api

EXPOSE 8080

CMD ["/neuralops-api"]
