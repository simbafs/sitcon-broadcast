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

	api := r.Group("/api")

	eventRepo := repository.NewEventInMemory()
	eventUsecase := usecase.NewEvent(eventRepo)
	ginrest.NewEvent(api.Group("/event"), eventUsecase)

	nowRepo := repository.NewNow()
	nowUsecase := usecase.NewNow(nowRepo)
	ginrest.NewNow(api.Group("/now"), nowUsecase)

	r.NoRoute(ginrest.NoRouteHandler())

	log.Fatal(r.Run(":3000"))
}
