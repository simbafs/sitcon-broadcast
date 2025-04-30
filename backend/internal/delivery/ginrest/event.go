package ginrest

import (
	"net/http"

	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Event struct {
	event usecase.Event
}

func NewEvent(r gin.IRouter, event usecase.Event) *Event {
	g := &Event{
		event: event,
	}

	r.POST("/event", g.Create)
	r.DELETE("/event/:name", g.Delete)
	r.GET("/event", g.List)
	r.GET("/event/:name", g.Get)
	r.PUT("/event/:name", g.UpdateScript)

	return g
}

// POST /event
func (g *Event) Create(c *gin.Context) {
	var input usecase.EventCreateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	output, err := g.event.Create(c, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// DELETE /event/{name}
func (g *Event) Delete(c *gin.Context) {
	name := c.Param("name")

	input := usecase.EventDeleteInput{
		Name: name,
	}

	output, err := g.event.Delete(c, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// GET /event/{name}
func (g *Event) Get(c *gin.Context) {
	name := c.Param("name")

	input := usecase.EventGetInput{
		Name: name,
	}

	output, err := g.event.Get(c, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// GEt /event
func (g *Event) List(c *gin.Context) {
	input := usecase.EventListInput{}

	output, err := g.event.List(c, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// PUT /event/{name} body: .script
func (g *Event) UpdateScript(c *gin.Context) {
	var input usecase.EventSetScriptInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	name := c.Param("name")
	input.Name = name

	output, err := g.event.Execute(c, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output)
}
