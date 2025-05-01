package ginrest

import (
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Counter struct {
	counter usecase.Counter
}

func NewCounter(r gin.IRouter, counter usecase.Counter) *Counter {
	c := &Counter{
		counter: counter,
	}

	r.POST("/", c.Create)
	r.GET("/", c.List)
	r.GET("/:name", c.Get)
	r.POST("/:name/start", c.Start)
	r.POST("/:name/stop", c.Stop)
	r.POST("/:name/setinit", c.SetInit)

	return c
}

type CounterCreateInput struct {
	Name string `json:"name" binding:"required"`
	Init int    `json:"init" binding:"required"`
}

// POST /counter
func (c *Counter) Create(ctx *gin.Context) {
	var input CounterCreateInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.Error(err)
		return
	}

	output, err := c.counter.New(input.Name, input.Init)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, output)
}

// GET /counter
func (c *Counter) List(ctx *gin.Context) {
	counters := c.counter.List()
	ctx.JSON(200, counters)
}

// GET /counter/{name}
func (c *Counter) Get(ctx *gin.Context) {
	name := ctx.Param("name")

	counter, err := c.counter.Get(name)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, counter)
}

// POST /counter/{name}/start
func (c *Counter) Start(ctx *gin.Context) {
	name := ctx.Param("name")

	err := c.counter.Start(name)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, nil)
}

// POST /counter/{name}/stop
func (c *Counter) Stop(ctx *gin.Context) {
	name := ctx.Param("name")

	err := c.counter.Stop(name)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, nil)
}

type CounterSetInitInput struct {
	Name string `json:"name" binding:"required"`
	Init int    `json:"init" binding:"required"`
}

// POST /counter/{name}/setinit
func (c *Counter) SetInit(ctx *gin.Context) {
	var input CounterSetInitInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = c.counter.SetInit(input.Name, input.Init)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, nil)
}
