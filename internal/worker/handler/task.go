package handler

import (
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/qmdx00/crobjob/internal/task/constant"
	"github.com/qmdx00/crobjob/internal/worker/job"
	"github.com/qmdx00/crobjob/internal/worker/server"
	"github.com/qmdx00/crobjob/rpc"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"strings"
)

// TaskHandler ...
type TaskHandler struct {
	cronMap     map[string]*cron.Cron
	scheduleMap map[string]cron.Schedule
	log         *zap.Logger
}

// NewTaskHandler ...
func NewTaskHandler(log *zap.Logger) server.Receive {
	return &TaskHandler{log: log, cronMap: make(map[string]*cron.Cron), scheduleMap: make(map[string]cron.Schedule)}
}

// Consume ...
func (t *TaskHandler) Consume(msg *sarama.ConsumerMessage) {
	// protobuf decode
	task := &rpc.Task_Model{}
	_ = proto.Unmarshal(msg.Value, task)

	switch strings.ToUpper(strings.TrimSpace(string(msg.Key))) {
	case constant.StartCommand:
		t.handleStart(task)
	case constant.StopCommand:
		t.handleStop(task)
	case constant.AddCommand:
		t.handleAdd(task)
	case constant.RemoveCommand:
		t.handleRemove(task)
	}
}

// handleAdd ...
func (t *TaskHandler) handleAdd(task *rpc.Task_Model) {
	t.log.Info("add task", zap.String("key", task.Key))

	schedule, err := cron.Parse(task.Expr)
	if err != nil {
		t.log.Error("parse expr error", zap.String("expr", task.Expr))
		return
	}

	cc := cron.New()
	cc.Schedule(schedule, job.NewTaskJob(t.log, task))

	t.cronMap[task.Key] = cc
	t.scheduleMap[task.Key] = schedule

}

// handleRemove ...
func (t *TaskHandler) handleRemove(task *rpc.Task_Model) {
	t.log.Info("remove task", zap.String("key", task.Key))
	if _, ok := t.cronMap[task.Key]; ok {
		t.handleStop(task)

		delete(t.cronMap, task.Key)
		delete(t.scheduleMap, task.Key)
	}
}

// handleStart ...
func (t *TaskHandler) handleStart(task *rpc.Task_Model) {
	t.log.Info("start task", zap.String("key", task.Key))
	if c, ok := t.cronMap[task.Key]; ok {
		c.Start()
	} else {
		t.handleAdd(task)
		if cc, exist := t.cronMap[task.Key]; exist {
			cc.Start()
		}
	}
}

// handleStop ...
func (t *TaskHandler) handleStop(task *rpc.Task_Model) {
	t.log.Info("stop task", zap.String("key", task.Key))
	if c, ok := t.cronMap[task.Key]; ok {
		c.Stop()
	}
}
