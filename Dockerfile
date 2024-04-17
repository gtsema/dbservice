FROM golang:alpine as builder
LABEL authors="gts"

COPY . .
RUN apk add --no-cache --update gcc g++ &&\
    cd ./cmd &&\
    CGO_ENABLED=1 go build -o dbservice

FROM alpine

COPY --from=builder go/cmd/dbservice ./dbservice/
COPY --from=builder go/configs ./dbservice/configs/
COPY --from=builder go/migrations ./dbservice/migrations/

RUN mkdir ./migrate &&\
    wget -O ./migrate/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz &&\
    tar -xzf ./migrate/migrate.tar.gz -C ./migrate &&\
    mv ./migrate/migrate /usr/local/bin/ &&\
    rm -r ./migrate &&\
    chmod +x /usr/local/bin/migrate
#    dbservice/dbservice
