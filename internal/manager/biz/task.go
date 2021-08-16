package biz

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func NewTaskBusiness(log *zap.Logger) *TaskBusiness {
	return &TaskBusiness{log: log}
}

type TaskBusiness struct {
	log *zap.Logger
}

func (*TaskBusiness) GetList(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
