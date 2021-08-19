package config

import (
	"github.com/qmdx00/crobjob/pkg/config"
	"github.com/spf13/viper"
)

type ManagerConfig struct {
	Viper *viper.Viper
}

func NewManagerConfig() *ManagerConfig {
	return &ManagerConfig{Viper: config.New()}
}
