FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o /usr/local/bin/workflow-runner ./cmd/workflow-runner

ENTRYPOINT ["/usr/local/bin/workflow-runner"]
