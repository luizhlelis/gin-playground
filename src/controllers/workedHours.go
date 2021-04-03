package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostWorkedHours(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, "Nice job!")
}
