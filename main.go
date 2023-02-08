package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"notifications/config"
	"notifications/handlers"
	r "notifications/routes"
)

func init() {
	//load values from .env
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

// @title           Документация сервиса уведомлений
// @version         1.0
// @description     Сервис уведомлений

// @contact.name   API Support

// @host      localhost
// @BasePath  /
// @schemes http
func main() {
	conf := config.GetConfig()
	routes := chi.NewRouter()

	logrus.Info("run server")

	// апи
	r.ApiRoutes(routes)

	// запуск сервера
	handlers.Error(http.ListenAndServe(conf.Server.Url, routes))
}
