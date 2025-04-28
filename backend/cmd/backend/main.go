package main

import (
	"log"

	"backend/internal/delivery/ginrest"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(ginrest.ErrorHandler())

	eventRepo := repository.NewEventInMemoryRepository()
	eventUsecase := usecase.NewEvent(eventRepo)
	handler := ginrest.New(eventUsecase)

	handler.Route(r.Group("/api"))
	r.NoRoute(ginrest.NoRouteHandler())

	log.Fatal(r.Run(":3000"))
}
