package server

import (
	"gin-example-structure/config"

	log "github.com/sirupsen/logrus"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	log.Info("Server Running On port ", config.GetString("server.port"))
	r.Run(config.GetString("server.port"))
}
