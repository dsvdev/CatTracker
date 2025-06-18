package config

import "os"

var AppConfig Config

func init() {
	AppConfig = Config{
		Db: DbConfig{URL: os.Getenv("DB_URL")},
		Kafka: KafkaConfig{
			URL:       os.Getenv("KAFKA_URL"),
			CatTopic:  os.Getenv("KAFKA_CAT_TOPIC"),
			JumpTopic: os.Getenv("KAFKA_JUMP_TOPIC"),
			FeedTopic: os.Getenv("KAFKA_FEED_TOPIC"),
		},
	}
}

type Config struct {
	Db    DbConfig
	Kafka KafkaConfig
}

type DbConfig struct {
	URL string
}

type KafkaConfig struct {
	URL       string
	CatTopic  string
	JumpTopic string
	FeedTopic string
}
