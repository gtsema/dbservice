package main

import (
	router "dbservice/pkg/transport/http"
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(),
	}

	fmt.Println("Запуск сервера на http://localhost:8080")
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
	}
}
