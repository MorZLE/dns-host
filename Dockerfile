FROM golang:latest

WORKDIR /app
COPY srv/go.mod go.sum ./

COPY ./srv/ ./


RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o srv/cmd/main.go

EXPOSE 44044
CMD ["sudo /dns-host"]