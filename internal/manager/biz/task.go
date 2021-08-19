package biz

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/qmdx00/crobjob/rpc"
	"go.uber.org/zap"
	"net/http"
)

// NewTaskBusiness ...
func NewTaskBusiness(log *zap.Logger, client rpc.TaskServiceClient, tracer opentracing.Tracer) *TaskBusiness {
	return &TaskBusiness{client: client, log: log, tracer: tracer}
}

// TaskBusiness ...
type TaskBusiness struct {
	client rpc.TaskServiceClient
	tracer opentracing.Tracer
	log    *zap.Logger
}

// CreateTask ...
func (b *TaskBusiness) CreateTask(ctx *gin.Context) {
	spanCtx, _ := ctx.Get("context")
	model := &rpc.Task_Model{
		Id:          0,
		Name:        "aaa",
		TaskType:    "aaa",
		TaskData:    "aaa",
		CronExpr:    "",
		Timeout:     0,
		Description: "aaa",
		CreatedAt:   "",
		UpdatedAt:   "",
	}

	// call rpc create task
	task, err := b.client.CreateTask(spanCtx.(context.Context), &rpc.Task_CreateTask{Data: model})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
