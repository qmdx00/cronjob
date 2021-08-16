package biz

import (
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

func (b *TaskBusiness) GetList(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
