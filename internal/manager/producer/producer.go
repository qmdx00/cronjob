package producer

import (
	"github.com/Shopify/sarama"
)

type TaskProducer struct {
	producer sarama.SyncProducer
}

func NewProducer() (*TaskProducer, error) {
	// HACK:
	brokers := []string{"127.0.0.1:9092"}

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &TaskProducer{producer: producer}, nil
}

func (t *TaskProducer) Send(key, value string) (int32, int64, error) {
	return t.producer.SendMessage(&sarama.ProducerMessage{
		// HACK:
		Topic: "test",
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	})
}
