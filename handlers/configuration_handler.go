package handlers

import (
	"nas-ars-go-projekat/models"
	"nas-ars-go-projekat/storage"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateConfiguration(c *gin.Context) {
	var config models.Configuration

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	config.ID = uuid.New().String()

	storage.Configurations[config.ID] = config

	c.JSON(http.StatusCreated, config)
}

func GetConfiguration(c *gin.Context) {
	name := c.Param("name")
	version := c.Param("version")

	for _, config := range storage.Configurations {
		if config.Name == name && config.Version == version {
			c.JSON(http.StatusOK, config)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Configuration not found",
	})
}

func GetConfigurationByID(c *gin.Context) {
	id := c.Param("id")

	config, exists := storage.Configurations[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Configuration not found",
		})
		return
	}

	c.JSON(http.StatusOK, config)
}

func DeleteConfiguration(c *gin.Context) {
	id := c.Param("id")

	_, exists := storage.Configurations[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Configuration not found",
		})
		return
	}

	delete(storage.Configurations, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Configuration deleted successfully",
	})
}

func DeleteConfigurationByVersion(c *gin.Context) {
	name := c.Param("name")
	version := c.Param("version")

	for id, config := range storage.Configurations {
		if config.Name == name && config.Version == version {
			delete(storage.Configurations, id)

			c.JSON(http.StatusOK, gin.H{
				"message": "Configuration deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Configuration not found",
	})
}

func ReplaceConfiguration(c *gin.Context) {
	id := c.Param("id")

	_, exists := storage.Configurations[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Configuration not found",
		})
		return
	}

	var newConfig models.Configuration

	if err := c.ShouldBindJSON(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	newConfig.ID = id

	storage.Configurations[id] = newConfig

	c.JSON(http.StatusOK, newConfig)
}
