package controllers

import (
	"gin-example-structure/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MatchController struct{}

var matchinfoSummaryModel = new(models.MatchinfoSummary)

// @Summary get matchinfo summary
// @Description get matchinfo by ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /match/{id} [get]
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
