package main

import (
	"context"
	"dbservice"
	"dbservice/internal/database"
	"dbservice/internal/repository"
	"dbservice/internal/service"
	http_ "dbservice/internal/transport/http"
	"dbservice/internal/transport/http/handler"
	"fmt"
	"github.com/spf13/viper"
	"log"
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

	// Инициализируем репу, серавис
	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	// Собираем строку подключения для сервера
	serverUrl := fmt.Sprintf("%s:%s", viper.GetString("server_host"), viper.GetString("server_port"))

	server := new(dbservice.Server)

	// Стааартуем
	go func() {
		if err := server.Run(serverUrl, http_.NewRouter(newHandler)); err != nil {
			log.Fatalf("failed to run http server: %s", err.Error())
		}
	}()

	log.Print("server started on ", serverUrl)

	// Посылаем в поток сервера привет
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("shutting down server... ")

	// Завершаем с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
