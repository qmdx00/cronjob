package config

import (
	"github.com/qmdx00/crobjob/pkg/config"
	"github.com/spf13/viper"
)

type WorkerConfig struct {
	Viper *viper.Viper
}

func NewWorkerConfig() *WorkerConfig {
	return &WorkerConfig{Viper: config.New()}
}
