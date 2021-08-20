package handler

import (
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/internal/worker/server"
	"go.uber.org/zap"
)

type TaskHandler struct {
	log *zap.Logger
}

func NewTaskHandler(log *zap.Logger) server.Receive {
	return &TaskHandler{log: log}
}

func (t TaskHandler) Consume(msg *sarama.ConsumerMessage) {
	t.log.Info("received task", zap.String(string(msg.Key), string(msg.Value)))
}
