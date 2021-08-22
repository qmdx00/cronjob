package job

import (
	"github.com/qmdx00/crobjob/rpc"
	"go.uber.org/zap"
)

// TaskJob ...
type TaskJob struct {
	log  *zap.Logger
	task *rpc.Task_Model
}

// NewTaskJob ...
func NewTaskJob(log *zap.Logger, task *rpc.Task_Model) *TaskJob {
	return &TaskJob{log: log, task: task}
}

// Run ...
func (t TaskJob) Run() {
	t.log.Info("run task", zap.String("task", t.task.Name))
}
