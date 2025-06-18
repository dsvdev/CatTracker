package db

import (
	"CatTracker/internal/event/model"
	"CatTracker/internal/general/db"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepository(client *db.PostgresClient) *EventRepo {
	return &EventRepo{db: client.Pool}
}

func (r *EventRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Event, error) {
	var event model.Event
	err := r.db.QueryRow(ctx, "SELECT id, cat_id, payload, action_type, created_at FROM actions WHERE id=$1", id).
		Scan(&event.ID, &event.CatID, &event.Payload, &event.CreatedAt)

	return &event, err
}

func (r *EventRepo) SaveNewEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	_, err := r.db.Exec(ctx,
		"INSERT INTO actions (id, cat_id, payload, action_type, created_at) VALUES ($1, $2, $3, $4, $5)",
		event.ID, event.CatID, event.Payload, event.Type, event.CreatedAt)
	return event, err
}
