package controllers

import (
	"gin-example-structure/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MatchController struct{}

var matchinfoSummaryModel = new(models.MatchinfoSummary)

func (u MatchController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		info, err := matchinfoSummaryModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve matchinfo", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "MatchinfoSummary founded!", "info": info})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
