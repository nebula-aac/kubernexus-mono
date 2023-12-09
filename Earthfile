VERSION 0.7
FROM golang:1.21-alpine
WORKDIR /workspace

deps:
    COPY go.mod go.sum ./
    RUN go mod vendor
    RUN go mod download
    SAVE ARTIFACT ./vendor
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

