package main

import (
	"log"

	"backend/internal/delivery/ginrest"
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/internal/usecase"
	"backend/sse"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(ginrest.ErrorHandler())

	api := r.Group("/api")

	sse := sse.New()

	eventRepo := repository.NewEventInMemory()
	eventUsecase := usecase.NewEvent(eventRepo)
	ginrest.NewEvent(api.Group("/event"), eventUsecase)

	nowRepo := entity.NewNow(0)
	nowUsecase := usecase.NewNow(nowRepo)
	ginrest.NewNow(api.Group("/now"), nowUsecase)

	counterRepo := repository.NewCounterInMemory(nil)
	counterUsecase := usecase.NewCounter(counterRepo, sse)
	ginrest.NewCounter(api.Group("/counter"), counterUsecase)

	r.NoRoute(ginrest.NoRouteHandler())

	go sse.Listen()

	log.Fatal(r.Run(":3000"))
}
