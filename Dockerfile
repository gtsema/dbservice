FROM golang:alpine as builder

ENV CGO_ENABLED=1

COPY . .
WORKDIR /go/cmd

RUN apk add --no-cache --update gcc g++ &&\
    go build -o dbservice &&\
    go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine

COPY --from=builder go/cmd/dbservice ./dbservice/
COPY --from=builder go/configs ./dbservice/configs/
COPY --from=builder go/migrations ./dbservice/migrations/
COPY --from=builder go/bin/goose ./dbservice/

RUN echo "mur"