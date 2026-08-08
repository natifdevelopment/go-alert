package alert

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

// Publisher sends alert payloads to a Kafka topic.
type Publisher struct {
	writer *kafka.Writer
	topic  string
}

// NewPublisher creates a Kafka publisher for the given brokers and topic.
func NewPublisher(brokers []string, topic string) *Publisher {
	return &Publisher{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			MaxAttempts:  3,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Async:        true,
		},
		topic: topic,
	}
}

// NewPublisherFromEnv creates a publisher using KAFKA_BROKERS and KAFKA_TOPIC env vars.
func NewPublisherFromEnv() (*Publisher, error) {
	brokers := []string{"localhost:9092"}
	if v := os.Getenv("KAFKA_BROKERS"); v != "" {
		brokers = []string{v}
	}
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "error-alerts"
	}
	return NewPublisher(brokers, topic), nil
}

// Publish sends the payload to Kafka asynchronously.
// Errors are returned only for synchronous serialization or write failures.
func (p *Publisher) Publish(ctx context.Context, payload Payload) error {
	value, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("alert: marshal payload: %w", err)
	}

	key := []byte(fmt.Sprintf("%s-%d", payload.Service, payload.Timestamp.UnixNano()))
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
		Time:  payload.Timestamp,
	})
}

// Close flushes and closes the Kafka writer.
func (p *Publisher) Close() error {
	return p.writer.Close()
}
