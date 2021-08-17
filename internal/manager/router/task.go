package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qmdx00/crobjob/internal/manager/biz"
)

// RegisterTaskRouter ...
func RegisterTaskRouter(group *gin.RouterGroup, task *biz.TaskBusiness) {
	group.GET("/tasks", task.CreateTask)
}
