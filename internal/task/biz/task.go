package biz

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/qmdx00/crobjob/rpc"
	"go.uber.org/zap"
)

var _ rpc.TaskServiceServer = (*TaskBusiness)(nil)

func NewTaskBusiness(log *zap.Logger) *TaskBusiness {
	return &TaskBusiness{log: log}
}

type TaskBusiness struct {
	log *zap.Logger
}

func (t TaskBusiness) GetListByType(_ context.Context, req *rpc.Task_GetListByType) (*rpc.Task_List, error) {
	return nil, nil
}

func (t TaskBusiness) CreateTask(ctx context.Context, req *rpc.Task_CreateTask) (*rpc.Task_Model, error) {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()

	return req.Data, nil
}
