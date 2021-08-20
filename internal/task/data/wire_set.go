package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/hoisie/mustache"
	"github.com/opentracing/opentracing-go"
	"github.com/qmdx00/crobjob/internal/task/config"
	"github.com/qmdx00/crobjob/pkg/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"
)

var ProviderSet = wire.NewSet(NewGormDB, NewTracer, NewTask)

func NewGormDB(tracer opentracing.Tracer, config *config.TaskConfig) (*gorm.DB, error) {
	url := mustache.Render("{{user}}:{{password}}@tcp({{endpoint}})/{{database}}?charset=utf8&parseTime=True&loc=Local", map[string]interface{}{
		"user":     config.Viper.GetString("resource.mysql.task.user"),
		"password": config.Viper.GetString("resource.mysql.task.password"),
		"database": config.Viper.GetString("resource.mysql.task.database"),
		"endpoint": config.Viper.GetString("resource.mysql.task.endpoint"),
	})

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.Use(gormopentracing.New(gormopentracing.WithTracer(tracer)))

	return db, nil
}

// NewTracer ...
func NewTracer(config *config.TaskConfig) (opentracing.Tracer, func(), error) {
	serviceName := config.Viper.GetString("task.log.prefix")
	agent := config.Viper.GetString("resource.jaeger.agent")
	tracer, closer, err := middleware.NewJaegerTracer(serviceName, agent)
	if err != nil {
		return nil, nil, err
	}
	return tracer, func() {
		_ = closer.Close()
	}, nil
}
