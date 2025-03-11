package main

import (
	"fmt"
	"online-music-library/config"
	"online-music-library/database"
	_ "online-music-library/docs"
	"online-music-library/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Онлайн-библиотека песен API
// @version 1.0
// @description Это REST API для управления музыкальной библиотекой
// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)

	// Инициализация маршрутизатора
	r := routes.SetupRouter()

	// Подключение Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
