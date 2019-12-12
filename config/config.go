package config

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config = viper.New()

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	log.WithField("FileName", env).Info("正在使用的配置文件")
	var err error
	config.SetConfigType("toml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("./")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		//viper配置发生变化了 执行响应的操作
		log.Info("Config file changed:", e.Name)
	})
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
