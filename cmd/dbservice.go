package main

import (
	"context"
	"dbservice/internal/database"
	"dbservice/internal/repository"
	"dbservice/internal/service"
	http_ "dbservice/internal/transport/http"
	"dbservice/internal/transport/http/handler"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// читаем конфиг
	if err := initCfg(); err != nil {
		log.Fatalf("Error reading configs file, %s", err.Error())
	}

	// Инициализируем подключение к бд sqlite
	db, err := database.NewSqliteDB(viper.GetString("db_url"))
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Инициализируем репу, сервис и хендлер
	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	serverUrl := fmt.Sprintf("%s:%s", viper.GetString("server_host"), viper.GetString("server_port"))

	// Создаём инстанс http-сервера
	server := new(http_.Server)

	// Стааартуем
	go func() {
		if err := server.Run(serverUrl, http_.NewRouter(newHandler)); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to run http server: %s", err.Error())
		}
	}()

	log.Print("server started on ", serverUrl)

	// Инициализируем канал для получения сигналов от операционной системы
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	// Блокируем до получения сигнала
	<-stopChan

	log.Print("shutting down server... ")

	// Создаем контекст с тайм-аутом для корректного завершения работы сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем Shutdown, чтобы завершить работу сервера, отработав все текущие запросы
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown server: %s", err.Error())
	}

	log.Print("server shutdown complete")

	// Завершаем работу с бд
	if err := db.Close(); err != nil {
		log.Fatalf("failed to close database: %s", err.Error())
	}
}

func initCfg() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
