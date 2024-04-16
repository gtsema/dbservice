package main

import (
	"dbservice/internal/repository"
	"dbservice/internal/service"
	router "dbservice/internal/transport/http"
	"dbservice/internal/transport/http/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {

	db, err := repository.NewSqliteDB(":memory:")
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(handler),
	}

	fmt.Println("Запуск сервера на http://localhost:8080")
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
	}
}
