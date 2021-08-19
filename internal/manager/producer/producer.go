package producer

import (
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/internal/manager/config"
)

type TaskProducer struct {
	producer sarama.SyncProducer
}

func NewProducer(config *config.ManagerConfig) (*TaskProducer, func(), error) {
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

func (t *TaskProducer) Send(topic, key, value string) (int32, int64, error) {
	return t.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	})
}
