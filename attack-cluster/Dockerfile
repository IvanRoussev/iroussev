FROM golang:1.24

WORKDIR /app

COPY attack-cluster/go.mod attack-cluster/go.sum ./

RUN go mod download

COPY attack-cluster/. ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./...

EXPOSE 8080

CMD ["/docker-gs-ping"]
