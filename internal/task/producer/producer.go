package producer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/qmdx00/crobjob/internal/task/config"
	"github.com/qmdx00/crobjob/internal/task/constant"
)

type TaskProducer struct {
	producer sarama.SyncProducer
	tracer   opentracing.Tracer
}

func NewProducer(config *config.TaskConfig) (*TaskProducer, func(), error) {
	brokers := config.Viper.GetStringSlice("resource.kafka.brokers")

	kafka := sarama.NewConfig()
	kafka.Producer.Retry.Max = config.Viper.GetInt("resource.kafka.retry.max")
	kafka.Producer.RequiredAcks = sarama.WaitForAll
	kafka.Producer.Return.Successes = true
	kafka.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(brokers, kafka)
	if err != nil {
		return nil, nil, err
	}

	return &TaskProducer{producer: producer}, func() {
		producer.Close()
	}, nil
}

func (t *TaskProducer) Send(ctx context.Context, key string, value []byte) (int32, int64, error) {
	span, _ := opentracing.StartSpanFromContext(
		ctx, key, opentracing.Tag{Key: string(ext.Component), Value: "Kafka producer"})
	defer span.Finish()

	return t.producer.SendMessage(&sarama.ProducerMessage{
		Topic: constant.TaskTopic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	})
}
