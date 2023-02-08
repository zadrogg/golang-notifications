package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"notifications/handlers"
	"time"
)

func ApiRoutes(r *chi.Mux) {
	r.Use(middleware.Timeout(30 * time.Second))

	r.Post("/{type}", handlers.SendNotification)
	r.Get("/docs/", httpSwagger.WrapHandler)
}
