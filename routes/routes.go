package routes

import (
	"nas-ars-go-projekat/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/configurations", handlers.CreateConfiguration)
	router.GET("/configurations/:name/:version", handlers.GetConfiguration)
	router.GET("/configurations/id/:id", handlers.GetConfigurationByID)
	router.DELETE("/configurations/id/:id", handlers.DeleteConfiguration)
	router.DELETE("/configurations/:name/:version", handlers.DeleteConfigurationByVersion)

	router.POST("/groups", handlers.CreateGroup)
	router.GET("/groups/:name/:version", handlers.GetGroup)
	router.GET("/groups/id/:id", handlers.GetGroupByID)
	router.DELETE("/groups/id/:id", handlers.DeleteGroup)
	router.DELETE("/groups/:name/:version", handlers.DeleteGroupByVersion)

	router.POST("/groups/id/:id/configurations/:configId", handlers.AddConfigurationToGroup)
	router.DELETE("/groups/id/:id/configurations/:configId", handlers.RemoveConfigurationFromGroup)
}
