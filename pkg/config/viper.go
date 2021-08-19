package config

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// New *viper.Viper config ...
func New() *viper.Viper {
	var conf = flag.String("config", "./runtime.config", "config file path")
	flag.Parse()

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(*conf)
	config.AddConfigPath("./")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("Read runtime config fail:", err.Error())
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	return config
}
