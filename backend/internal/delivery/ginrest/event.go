package ginrest

import (
	"net/http"

	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// POST /event
func (g *Gin) Create(c *gin.Context) {
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
func (g *Gin) Delete(c *gin.Context) {
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
func (g *Gin) Get(c *gin.Context) {
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
func (g *Gin) List(c *gin.Context) {
	input := usecase.EventListInput{}

	output, err := g.event.List(c, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// PUT /event/{name} body: .script
func (g *Gin) UpdateScript(c *gin.Context) {
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
