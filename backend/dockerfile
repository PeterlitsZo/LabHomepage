FROM golang:alpine

ARG GOPROXY

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy && go build -v -o /usr/local/bin/backend .
