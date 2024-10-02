package controllers

import (
	"github.com/gatij/gin-golang-app/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetData(c *gin.Context) {
	data := services.GetDataService()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
