package biz

import (
	"context"
	"github.com/qmdx00/crobjob/internal/task/data"
	"github.com/qmdx00/crobjob/internal/task/producer"
	"github.com/qmdx00/crobjob/rpc"
	"go.uber.org/zap"
)

var _ rpc.TaskServiceServer = (*TaskBusiness)(nil)

// NewTaskBusiness ...
func NewTaskBusiness(log *zap.Logger, producer *producer.TaskProducer, task data.TaskRepo) *TaskBusiness {
	return &TaskBusiness{log: log, producer: producer, task: task}
}

// TaskBusiness ...
type TaskBusiness struct {
	producer *producer.TaskProducer
	log      *zap.Logger
	task     data.TaskRepo
}

// GetListByType ...
func (t TaskBusiness) GetListByType(ctx context.Context, req *rpc.Task_GetListByType) (*rpc.Task_List, error) {
	return nil, nil
}

// CreateTask ...
func (t TaskBusiness) CreateTask(ctx context.Context, req *rpc.Task_CreateTask) (*rpc.Task_Model, error) {
	t.producer.Send(ctx, "hello", "world")
	task, _ := t.task.CreateTask(ctx, req.Data)
	return task, nil
}
