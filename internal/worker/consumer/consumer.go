package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/internal/task/constant"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/pkg/transport"
	"go.uber.org/zap"
)

type TaskConsumer struct {
	consumer sarama.Consumer
	log      *zap.Logger
	//message  chan<- *sarama.ConsumerMessage
}

func NewTaskConsumer(config *config.WorkerConfig, log *zap.Logger) (transport.Server, error) {
	brokers := config.Viper.GetStringSlice("resource.kafka.brokers")
	consumer, err := sarama.NewConsumer(brokers, sarama.NewConfig())
	if err != nil {
		return nil, err
	}

	return &TaskConsumer{consumer: consumer, log: log}, nil
}

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
			//t.message <- msg
			// TODO 抽象消费方法
			t.log.Info("received message", zap.String(string(msg.Key), string(msg.Value)))
		}
	}
}

func (t *TaskConsumer) Stop(ctx context.Context) error {
	return t.consumer.Close()
}
