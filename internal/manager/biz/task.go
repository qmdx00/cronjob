package biz

import (
	"context"
	"errors"
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
	task := &rpc.Task_Model{}

	if err := ctx.BindJSON(task); err != nil {
		ctx.Error(err)
		return
	}

	spanCtx, _ := ctx.Get("context")
	created, err := b.client.CreateTask(spanCtx.(context.Context), &rpc.Task_CreateTask{Data: task})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": created,
	})
}

// DeleteTask ...
func (b *TaskBusiness) DeleteTask(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.Error(errors.New("params error"))
		return
	}

	spanCtx, _ := ctx.Get("context")
	deleted, err := b.client.DeleteTask(spanCtx.(context.Context), &rpc.Task_DeleteTask{Key: key})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": deleted,
	})
}

func (b *TaskBusiness) GetAllTask(ctx *gin.Context) {
	spanCtx, _ := ctx.Get("context")

	list, err := b.client.GetAllTask(spanCtx.(context.Context), &rpc.Task_GetAllTask{})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tasks": list,
	})
}

func (b *TaskBusiness) GetTaskByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.Error(errors.New("params error"))
		return
	}

	spanCtx, _ := ctx.Get("context")
	task, err := b.client.GetByTaskId(spanCtx.(context.Context), &rpc.Task_GetTaskByKey{Key: key})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (b *TaskBusiness) StartTask(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.Error(errors.New("params error"))
		return
	}

	spanCtx, _ := ctx.Get("context")
	status, err := b.client.StartTask(spanCtx.(context.Context), &rpc.Task_StartTask{Key: key})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": status.Message,
	})
}

func (b *TaskBusiness) StopTask(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.Error(errors.New("params error"))
		return
	}

	spanCtx, _ := ctx.Get("context")
	status, err := b.client.StopTask(spanCtx.(context.Context), &rpc.Task_StopTask{Key: key})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": status.Message,
	})
}
