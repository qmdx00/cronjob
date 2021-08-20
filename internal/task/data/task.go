package data

import (
	"context"
	"errors"
	"github.com/qmdx00/crobjob/rpc"
	"gorm.io/gorm"
)

// TaskRepo ...
type TaskRepo interface {
	CreateTask(context.Context, *rpc.Task_CreateTask) (*rpc.Task_Model, error)
	GetByTaskId(context.Context, *rpc.Task_GetTaskByKey) (*rpc.Task_Model, error)
	GetAllTask(context.Context, *rpc.Task_GetAllTask) (*rpc.Task_List, error)
}

var _ TaskRepo = (*Task)(nil)

// Task ...
type Task struct {
	db *gorm.DB
}

// NewTask ...
func NewTask(db *gorm.DB) TaskRepo {
	return &Task{db: db}
}

// CreateTask ...
func (t *Task) CreateTask(ctx context.Context, req *rpc.Task_CreateTask) (*rpc.Task_Model, error) {
	task := &rpc.Task_Model{
		Key:         req.Data.Key,
		Name:        req.Data.Name,
		TaskType:    req.Data.TaskType,
		Expr:        req.Data.Expr,
		Description: req.Data.Description,
		Extra:       req.Data.Extra,
	}
	r := t.db.WithContext(ctx).Model(rpc.Task_Model{}).Create(task)
	return task, r.Error
}

// GetByTaskId ...
func (t *Task) GetByTaskId(ctx context.Context, req *rpc.Task_GetTaskByKey) (*rpc.Task_Model, error) {
	task := &rpc.Task_Model{}
	r := t.db.WithContext(ctx).Model(rpc.Task_Model{}).Where("`key` = ?", req.Key).First(task)
	if r.Error != nil && !errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	return task, nil
}

// GetAllTask ...
func (t *Task) GetAllTask(ctx context.Context, _ *rpc.Task_GetAllTask) (*rpc.Task_List, error) {
	tasks := make([]*rpc.Task_Model, 0)
	r := t.db.WithContext(ctx).Model(rpc.Task_Model{}).Find(&tasks)
	return &rpc.Task_List{List: tasks}, r.Error
}
