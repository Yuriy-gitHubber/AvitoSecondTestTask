package main

import (
	"log"
	"net/http"

	"ZADANIE-6105/config"
	"ZADANIE-6105/controllers"
	"ZADANIE-6105/models"
	"ZADANIE-6105/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загрузка переменных окружения
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}

	// Инициализация конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.PostgresConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Запуск миграций
	err = models.RunMigrations(db)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Инициализация контроллеров с подключением к базе данных
	controllers.InitControllers(db)

	// Инициализация маршрутов
	router := routes.InitRoutes()

	// Запуск HTTP-сервера
	log.Printf("Server is running on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
