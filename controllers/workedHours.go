package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func postWorkedHours(ginContext *gin.Context)
{
	ginContext.String(http.StatusOK, "Nice job!")
}