package db

import (
	"CatTracker/internal/cat/model"
	"CatTracker/internal/general/db"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CatRepository struct {
	db *pgxpool.Pool
}

func NewCatRepository(client *db.PostgresClient) *CatRepository {
	return &CatRepository{db: client.Pool}
}

func (r *CatRepository) GetByID(ctx context.Context, id int) (*model.Cat, error) {
	var cat model.Cat
	err := r.db.QueryRow(ctx, "SELECT id, color, name, age FROM cats WHERE id=$1", id).
		Scan(&cat.ID, &cat.Color, &cat.Name, &cat.Age)

	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *CatRepository) SaveNewCat(ctx context.Context, cat *model.Cat) (*model.Cat, error) {
	_, err := r.db.Exec(ctx, "INSERT INTO public.cats (id, name, age, color) VALUES ($1, $2, $3, $4)", cat.ID, cat.Name, cat.Age, cat.Color)
	return cat, err
}
