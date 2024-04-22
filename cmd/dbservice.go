package main

import (
	"dbservice/internal/repository"
	"dbservice/internal/service"
	router "dbservice/internal/transport/http"
	"dbservice/internal/transport/http/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	if err := initCfg(); err != nil {
		log.Fatalf("Error reading configs file, %s", err.Error())
	}

	db, err := repository.NewSqliteDB(viper.GetString("db_url"))
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository)
	userHandler := handler.NewHandler(userService)

	serverConStr := fmt.Sprintf("%s:%s", viper.GetString("server_host"), viper.GetString("server_port"))
	httpServer := &http.Server{
		Addr:    serverConStr,
		Handler: router.NewRouter(userHandler),
	}

	fmt.Println("Запуск сервера на " + serverConStr)
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
	}
}

func initCfg() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
