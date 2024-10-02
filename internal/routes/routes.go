package routes

import (
	"github.com/gatij/gin-golang-app/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)
	router.GET("/data", controllers.GetData)
}
