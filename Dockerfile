# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY domain ./domain
COPY filters ./filters 
COPY proxy ./proxy
COPY replication ./replication
COPY consumer ./consumer

RUN go build -o /shadowd ./cmd/shadowd

EXPOSE 8080

CMD [ "/shadowd" ]
