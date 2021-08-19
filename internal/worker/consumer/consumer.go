package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/pkg/transport"
)

type TaskConsumer struct {
	consumer sarama.Consumer
}

func NewTaskConsumer(config *config.WorkerConfig) (transport.Server, error) {
	brokers := config.Viper.GetStringSlice("resource.kafka.brokers")

	kafka := sarama.NewConfig()
	kafka.Producer.Retry.Max = 5
	kafka.Producer.RequiredAcks = sarama.WaitForAll

	consumer, err := sarama.NewConsumer(brokers, kafka)
	if err != nil {
		return nil, err
	}

	return &TaskConsumer{consumer: consumer}, nil
}

func (t *TaskConsumer) Start(ctx context.Context) error {
	// HACK:
	part, err := t.consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-part.Messages():
			// TODO 抽象消费方法
			fmt.Println(string(msg.Key), string(msg.Value))
		}
	}
}

func (t *TaskConsumer) Stop(ctx context.Context) error {
	return t.consumer.Close()
}
