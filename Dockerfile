# Use the official Ubuntu image
FROM golang:alpine as builder

WORKDIR /app

COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN go build -o server srv/cmd/main.go



FROM ubuntu:latest as Worker

COPY --from=builder /app /app

EXPOSE 44044

WORKDIR /app

RUN  apt-get update \
        &&  apt-get upgrade -y

CMD ["./server"]