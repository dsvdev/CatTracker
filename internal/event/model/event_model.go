package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type EventType string

const (
	Feed EventType = "Feed"
	Jump EventType = "Jump"
)

type Event struct {
	ID        uuid.UUID       `json:"id"`
	CatID     uuid.UUID       `json:"cat_id"`
	Payload   json.RawMessage `json:"payload"`
	Type      EventType       `json:"event_type"`
	CreatedAt time.Time       `json:"created_at"`
}

func NewEvent(req *NewEventRequest) *Event {
	return &Event{
		ID:        uuid.New(),
		CatID:     req.CatID,
		Payload:   req.Payload,
		Type:      EventType(req.Type),
		CreatedAt: time.Now(),
	}
}

type NewEventRequest struct {
	CatID   uuid.UUID       `json:"cat_id" binding:"required"`
	Payload json.RawMessage `json:"payload" binding:"required"`
	Type    string          `json:"event_type" binding:"required,oneof=Feed Jump"`
}
