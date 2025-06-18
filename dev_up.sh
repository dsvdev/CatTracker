#!/bin/bash

set -e

echo "🚀 Запуск docker-compose..."
docker-compose up -d

echo "⏳ Ожидание PostgreSQL..."
until docker exec cat_postgres pg_isready -U cat_user > /dev/null 2>&1; do
  sleep 1
done

echo "✅ PostgreSQL готов"

echo "📦 Применение миграций..."
# Замени путь на свой бин мигратора, если используешь goose или migrate
migrate -path migrations -database "postgres://cat_user:cat_pass@localhost:5432/cat_db?sslmode=disable" up

echo "⏳ Ожидание Kafka..."
until nc -z localhost 9092; do
  sleep 1
done

echo "✅ Kafka готов"

echo "🎯 Создание топиков..."

docker exec kafka kafka-topics.sh \
  --create --if-not-exists \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 --partitions 3 \
  --topic cat.general

docker exec kafka kafka-topics.sh \
  --create --if-not-exists \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 --partitions 3 \
  --topic cat.feed

docker exec kafka kafka-topics.sh \
  --create --if-not-exists \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 --partitions 3 \
  --topic cat.jumps

echo "✅ Всё готово к работе!"
