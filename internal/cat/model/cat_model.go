package model

import (
	uuid "github.com/google/uuid"
)

type Cat struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Color string    `json:"color"`
	Age   int       `json:"age"`
}

type CreateCatRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required,oneof=black orange brown"`
	Age   int    `json:"age" binding:"required,gte=1,lte=15"`
}
