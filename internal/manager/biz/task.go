package biz

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmdx00/crobjob/rpc"
	"go.uber.org/zap"
	"net/http"
)

func NewTaskBusiness(log *zap.Logger, client rpc.TaskServiceClient) *TaskBusiness {
	return &TaskBusiness{client: client, log: log}
}

type TaskBusiness struct {
	client rpc.TaskServiceClient
	log    *zap.Logger
}

func (b *TaskBusiness) CreateTask(ctx *gin.Context) {
	task, err := b.client.CreateTask(context.Background(), &rpc.Task_CreateTask{Data: &rpc.Task_Model{
		Id:          0,
		Name:        "aaa",
		TaskType:    "aaa",
		TaskData:    "aaa",
		CronExpr:    "",
		Timeout:     0,
		Description: "aaa",
		CreatedAt:   "",
		UpdatedAt:   "",
	}})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
