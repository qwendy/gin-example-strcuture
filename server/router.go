package server

import (
	"gin-example-structure/controllers"
	"gin-example-structure/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	hello := new(controllers.HelloController)

	router.GET("/hello", hello.Status)
	router.Use(middlewares.AuthMiddleware())

	matchGroup := router.Group("match")
	{
		matchinfoSummary := new(controllers.MatchController)
		matchGroup.GET("/:id", matchinfoSummary.Retrieve)
	}
	return router

}
