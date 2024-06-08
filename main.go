package main

import (
	_ "autoservice-backend/docs" // Этот импорт важен для Swagger документации
	"autoservice-backend/internal"
	"autoservice-backend/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Подключение к базе данных
	internal.ConnectDatabase()

	// Настройка маршрутов
	router := routes.SetupRouter()

	// Добавление Swagger маршрута
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера на порту 8080
	router.Run(":8080")
}
