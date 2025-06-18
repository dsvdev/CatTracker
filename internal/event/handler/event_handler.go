package handler

import (
	"CatTracker/internal/event/model"
	"CatTracker/internal/event/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewEventHandler(eventService *service.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

type EventHandler struct {
	eventService *service.EventService
}

func (h *EventHandler) RegisterRoutes(g *gin.Engine) {
	g.POST("/event", h.createEvent)
}

func (h *EventHandler) createEvent(c *gin.Context) {
	var event model.NewEventRequest
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	newEvent := model.NewEvent(&event)

	if err := h.eventService.ProcessNewEvent(context.Background(), newEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}
