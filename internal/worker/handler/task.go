package handler

import (
	"github.com/Shopify/sarama"
	"github.com/qmdx00/crobjob/internal/task/constant"
	"github.com/qmdx00/crobjob/internal/worker/server"
	"go.uber.org/zap"
	"strings"
)

type TaskHandler struct {
	log *zap.Logger
}

func NewTaskHandler(log *zap.Logger) server.Receive {
	return &TaskHandler{log: log}
}

func (t *TaskHandler) Consume(msg *sarama.ConsumerMessage) {
	switch strings.ToUpper(strings.TrimSpace(string(msg.Value))) {
	case constant.StartCommand:
		t.handleStart(msg)
	case constant.StopCommand:
		t.handleStop(msg)
	}
}

func (t *TaskHandler) handleStart(msg *sarama.ConsumerMessage) {
	t.log.Info("received task", zap.String(string(msg.Key), string(msg.Value)))

}

func (t *TaskHandler) handleStop(msg *sarama.ConsumerMessage) {
	t.log.Info("received task", zap.String(string(msg.Key), string(msg.Value)))

}
