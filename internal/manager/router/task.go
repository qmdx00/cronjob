package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qmdx00/crobjob/internal/manager/biz"
)

// RegisterTaskRouter ...
func RegisterTaskRouter(group *gin.RouterGroup, task *biz.TaskBusiness) {
	group.POST("/tasks", task.CreateTask)
	group.GET("/tasks", task.GetAllTask)
	group.GET("/tasks/:key", task.GetTaskByKey)
	group.GET("/tasks/:key/actions/start", task.StartTask)
	group.GET("/tasks/:key/actions/stop", task.StopTask)

}
