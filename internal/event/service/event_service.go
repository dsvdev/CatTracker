package service

import (
	"CatTracker/internal/event/db"
	"CatTracker/internal/event/model"
	"CatTracker/internal/general/kafka"
	"context"
)

func NewEventService(kafkaClient *kafka.KafkaClient, eventRepo *db.EventRepo) *EventService {
	return &EventService{
		kafkaClient: kafkaClient,
		eventRepo:   eventRepo,
	}
}

type EventService struct {
	kafkaClient *kafka.KafkaClient
	eventRepo   *db.EventRepo
}

func (s *EventService) ProcessNewEvent(ctx context.Context, event *model.Event) error {
	_, err := s.eventRepo.SaveNewEvent(ctx, event)
	if err != nil {
		return err
	}
	err = s.kafkaClient.SendNewEvent(event)
	return err
}
