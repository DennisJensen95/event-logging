VERSION 0.6

# Golang applications
golang-files-env:
    FROM golang:latest
    WORKDIR /app
    COPY cmd cmd
    COPY go.mod .
    COPY go.sum .

build-cul-linux-x86-64:
    FROM +golang-files-env
    RUN go build -o bin/cul ./cmd/computer-utilization-logging
    SAVE ARTIFACT ./bin/cul AS LOCAL ./bin/cul

unit-test-golang:
    FROM +golang-files-env
    RUN go test ./cmd/computer-utilization-logging

# Frontend applications
frontend-files-env:
    FROM node:16-alpine
    WORKDIR /app
    COPY ./frontend .

unit-test-frontend:
    FROM +frontend-files-env
    RUN CI=true npm install
    RUN CI=true npm test