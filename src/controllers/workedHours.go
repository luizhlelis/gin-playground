package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhlelis/gin-playground/models"
)

//PostWorkedHours is an API method to an employee register hours
func PostWorkedHours(ginContext *gin.Context) {
	var json TimeEntry

	if err := ginContext.ShouldBindJSON(&json); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, "Nice job!")
}
