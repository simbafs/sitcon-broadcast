package main

import (
	"log"

	"backend/internal/delivery/ginrest"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.Use(ginrest.ErrorHandler())

	eventRepo := repository.NewEventInMemoryRepository()
	eventUsecase := usecase.NewEvent(eventRepo)
	handler := ginrest.New(eventUsecase)

	handler.Route(g.Group("/api"))

	log.Fatal(g.Run(":3000"))
}
