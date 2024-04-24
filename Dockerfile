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
COPY --from=builder go/wait_database.sh ./dbservice/
COPY --from=builder go/migrations ./dbservice/migrations/
COPY --from=builder go/bin/goose ./dbservice/

WORKDIR ./dbservice

ENV GOOSE_DRIVER=sqlite3
ENV GOOSE_DBSTRING=./dbservice.db
ENV GOOSE_MIGRATION_DIR=./migrations

RUN chmod +x ./wait_database.sh

CMD ["./wait_database.sh"]
