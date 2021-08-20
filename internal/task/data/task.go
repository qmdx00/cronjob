package data

import (
	"context"
	"github.com/qmdx00/crobjob/rpc"
	"gorm.io/gorm"
)

type TaskRepo interface {
	CreateTask(context.Context, *rpc.Task_Model) (*rpc.Task_Model, error)
}

type Task struct {
	db *gorm.DB
}

func NewTask(db *gorm.DB) TaskRepo {
	return &Task{db: db}
}

func (t *Task) CreateTask(ctx context.Context, task *rpc.Task_Model) (*rpc.Task_Model, error) {
	created := &rpc.Task_Model{
		Name:        task.Name,
		TaskType:    task.TaskType,
		Description: task.Description,
	}
	r := t.db.WithContext(ctx).Model(rpc.Task_Model{}).Create(created)
	return created, r.Error
}
