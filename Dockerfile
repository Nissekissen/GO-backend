FROM golang:latest

WORKDIR /usr/src/app

COPY backend .
RUN go mod tidy
