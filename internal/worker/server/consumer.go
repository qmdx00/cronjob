package server

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/internal/task/constant"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/pkg/transport"
	"go.uber.org/zap"
)

// Receive ...
type Receive interface {
	Consume(msg *sarama.ConsumerMessage)
}

// TaskConsumer ...
type TaskConsumer struct {
	consumer sarama.Consumer
	log      *zap.Logger
	receive  Receive
}

// NewTaskConsumer ...
func NewTaskConsumer(config *config.WorkerConfig, log *zap.Logger, receive Receive) (transport.Server, error) {
	brokers := config.Viper.GetStringSlice("resource.kafka.brokers")
	consumer, err := sarama.NewConsumer(brokers, sarama.NewConfig())
	if err != nil {
		return nil, err
	}

	return &TaskConsumer{consumer: consumer, log: log, receive: receive}, nil
}

// Start ...
func (t *TaskConsumer) Start(ctx context.Context) error {
	part, err := t.consumer.ConsumePartition(constant.TaskTopic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-part.Messages():
			t.receive.Consume(msg)
		}
	}
}

// Stop ...
func (t *TaskConsumer) Stop(_ context.Context) error {
	return t.consumer.Close()
}
