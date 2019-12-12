package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct{}

func (h HelloController) Status(c *gin.Context) {
	c.String(http.StatusOK, "hi!")
}
