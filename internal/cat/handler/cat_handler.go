package handler

import (
	"CatTracker/internal/cat/model"
	"CatTracker/internal/cat/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func NewCatHandler(catService *service.CatService) *CatHandler {
	return &CatHandler{
		catService: catService,
	}
}

type CatHandler struct {
	catService *service.CatService
}

func (h *CatHandler) RegisterRoutes(g *gin.Engine) {
	g.POST("/cat", h.createCat)
}

func (h *CatHandler) createCat(c *gin.Context) {
	var cat model.CreateCatRequest
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	newCat := model.Cat{
		ID:    uuid.New(),
		Name:  cat.Name,
		Color: cat.Color,
		Age:   cat.Age,
	}

	if err := h.catService.ProcessNewCat(context.Background(), &newCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCat)
}
