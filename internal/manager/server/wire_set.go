package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/qmdx00/crobjob/internal/manager/biz"
	"github.com/qmdx00/crobjob/internal/manager/router"
	"github.com/qmdx00/crobjob/pkg/middleware"
	"go.uber.org/zap"
)

// ProviderSet for server ...
var ProviderSet = wire.NewSet(NewHttpServer, NewHTTPRouter)

// NewHTTPRouter ...
func NewHTTPRouter(log *zap.Logger, task *biz.TaskBusiness, tracer opentracing.Tracer) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	// use middlewares
	engine.Use(middleware.Ginzap(log))
	engine.Use(middleware.RecoveryWithZap(log, true))
	engine.Use(middleware.Opentracing(tracer))

	// set router
	mng := engine.Group("/v1/mng")
	router.RegisterTaskRouter(mng, task)

	return engine
}
