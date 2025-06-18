package kafka

import (
	catmodel "CatTracker/internal/cat/model"
	eventmodel "CatTracker/internal/event/model"
	"CatTracker/internal/general/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaClient(kafkaConfig *config.KafkaConfig) (*KafkaClient, error) {
	cl := &KafkaClient{
		conf: kafkaConfig,
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cl.conf.URL,
	})
	if err != nil {
		return nil, err
	}

	cl.p = p

	return cl, nil
}

type KafkaClient struct {
	conf *config.KafkaConfig
	p    *kafka.Producer
}

func (c *KafkaClient) SendNewCat(cat *catmodel.Cat) error {
	data, err := json.Marshal(cat)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	err = c.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &c.conf.CatTopic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	return err
}

func (c *KafkaClient) SendNewEvent(event *eventmodel.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	var topic string
	switch event.Type {
	case eventmodel.Jump:
		topic = c.conf.JumpTopic
	case eventmodel.Feed:
		topic = c.conf.FeedTopic
	default:
		return errors.New(fmt.Sprintf("unknown event type %s", event.Type))
	}

	err = c.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	return err
}

func (c *KafkaClient) Close() {
	c.p.Flush(1500)
	c.p.Close()
}
