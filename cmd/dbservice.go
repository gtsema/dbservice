package main

import (
	"dbservice/internal/database"
	"dbservice/internal/repository"
	"dbservice/internal/service"
	router "dbservice/internal/transport/http"
	"dbservice/internal/transport/http/handler"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initCfg(); err != nil {
		log.Fatalf("Error reading configs file, %s", err.Error())
	}

	db, err := database.NewSqliteDB(viper.GetString("db_url"))
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
	//exePath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//exeDir := filepath.Dir(exePath)
	//viper.AddConfigPath(exeDir + string(os.PathSeparator) + "configs")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
