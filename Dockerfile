# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY modules/ ./modules
COPY *.go ./

RUN go build -o /go-utils-service

EXPOSE 8080

CMD [ "/go-utils-service" ]
