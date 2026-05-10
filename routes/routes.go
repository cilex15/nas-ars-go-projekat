package routes

import (
	"nas-ars-go-projekat/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/configurations", handlers.CreateConfiguration)
	router.GET("/configurations/:name/:version", handlers.GetConfiguration)
	router.GET("/configurations/id/:id", handlers.GetConfigurationByID)
	router.DELETE("/configurations/:id", handlers.DeleteConfiguration)
}
