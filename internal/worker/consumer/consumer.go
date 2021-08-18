package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/pkg/transport"
)

type TaskConsumer struct {
	consumer sarama.Consumer
}

func NewTaskConsumer() (transport.Server, error) {
	// HACK:
	brokers := []string{"127.0.0.1:9092"}

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &TaskConsumer{consumer: consumer}, nil
}

func (t *TaskConsumer) Start(ctx context.Context) error {
	// HACK:
	part, err := t.consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-part.Messages():
			// TODO
			fmt.Println(string(msg.Key), string(msg.Value))
		}
	}
}

func (t *TaskConsumer) Stop(ctx context.Context) error {
	return t.consumer.Close()
}
