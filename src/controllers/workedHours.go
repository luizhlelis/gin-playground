package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhlelis/gin-playground/models"
)

func binder(objectToBind interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindQuery(&objectToBind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}

//PostWorkedHours is an API method to an employee register hours
func PostWorkedHours(ginContext *gin.Context) {
	var json models.TimeEntry

	// this aproach
	binder(json)

	// is equivant to this
	if err := ginContext.ShouldBindJSON(&json); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, "Nice job!")
}
