package config

import (
	"github.com/qmdx00/crobjob/pkg/config"
	"github.com/spf13/viper"
)

type TaskConfig struct {
	Viper *viper.Viper
}

func NewTaskConfig() *TaskConfig {
	return &TaskConfig{Viper: config.New()}
}
