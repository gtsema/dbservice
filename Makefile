BINARY_NAME=dbservice

build:
	docker-compose build ${BINARY_NAME}

run:
	docker-compose up ${BINARY_NAME}

migrate:
	GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./dbservice.db GOOSE_MIGRATION_DIR=./migrations goose up

#curl -d '{"chatId": "123", "name": "gosha"}' -H "application/json" -X POST 127.0.0.1:8080/user
#.\goose.exe -dir .\migrations\ sqlite3 .\dbservice.db up