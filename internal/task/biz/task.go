package biz

import (
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/qmdx00/crobjob/internal/task/constant"
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

func (t *TaskBusiness) GetByTaskId(ctx context.Context, req *rpc.Task_GetTaskByKey) (*rpc.Task_Model, error) {
	if req.Key == "" {
		return nil, errors.New("params error")
	}
	task, err := t.task.GetByTaskId(ctx, req)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskBusiness) GetAllTask(ctx context.Context, req *rpc.Task_GetAllTask) (*rpc.Task_List, error) {
	list, err := t.task.GetAllTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// CreateTask ...
func (t *TaskBusiness) CreateTask(ctx context.Context, req *rpc.Task_CreateTask) (*rpc.Task_Model, error) {
	if req.Data == nil {
		return nil, errors.New("bad request")
	}
	task, _ := t.task.CreateTask(ctx, req)
	// protobuf encode
	binary, _ := proto.Marshal(task)
	t.producer.Send(ctx, constant.AddCommand, binary)
	return task, nil
}

func (t *TaskBusiness) DeleteTask(ctx context.Context, req *rpc.Task_DeleteTask) (*rpc.Task_Model, error) {
	if req.Key == "" {
		return nil, errors.New("params error")
	}
	task, _ := t.task.DeleteTask(ctx, req)
	// protobuf encode
	binary, _ := proto.Marshal(task)
	t.producer.Send(ctx, constant.RemoveCommand, binary)
	return task, nil
}

// StartTask ...
func (t *TaskBusiness) StartTask(ctx context.Context, req *rpc.Task_StartTask) (*rpc.Task_StartTaskResp, error) {
	if req.Key == "" {
		return nil, errors.New("params error")
	}
	task, err := t.task.GetByTaskId(ctx, &rpc.Task_GetTaskByKey{Key: req.Key})
	if err != nil {
		return nil, err
	}
	// protobuf encode
	binary, _ := proto.Marshal(task)
	t.producer.Send(ctx, constant.StartCommand, binary)

	return &rpc.Task_StartTaskResp{Message: "START SEND"}, nil
}

// StopTask ...
func (t TaskBusiness) StopTask(ctx context.Context, req *rpc.Task_StopTask) (*rpc.Task_StopTaskResp, error) {
	if req.Key == "" {
		return nil, errors.New("params error")
	}
	task, err := t.task.GetByTaskId(ctx, &rpc.Task_GetTaskByKey{Key: req.Key})
	if err != nil {
		return nil, err
	}
	// protobuf encode
	binary, _ := proto.Marshal(task)
	t.producer.Send(ctx, constant.StopCommand, binary)

	return &rpc.Task_StopTaskResp{Message: "STOP SEND"}, nil
}
