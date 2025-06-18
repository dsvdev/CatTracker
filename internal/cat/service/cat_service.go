package service

import (
	"CatTracker/internal/cat/db"
	"CatTracker/internal/cat/model"
	"CatTracker/internal/general/kafka"
	"context"
)

func NewCatService(kafkaClient *kafka.KafkaClient, catRepo *db.CatRepository) *CatService {
	return &CatService{
		kafkaClient: kafkaClient,
		catRepo:     catRepo,
	}
}

type CatService struct {
	kafkaClient *kafka.KafkaClient
	catRepo     *db.CatRepository
}

func (s *CatService) ProcessNewCat(ctx context.Context, cat *model.Cat) error {
	_, err := s.catRepo.SaveNewCat(ctx, cat)
	if err != nil {
		return err
	}
	err = s.kafkaClient.SendNewCat(cat)
	return err
}
