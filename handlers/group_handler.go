package handlers

import (
	"nas-ars-go-projekat/models"
	"nas-ars-go-projekat/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateGroup(c *gin.Context) {
	var group models.ConfigurationGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	group.ID = uuid.New().String()

	storage.Groups[group.ID] = group

	c.JSON(http.StatusCreated, group)
}

func GetGroup(c *gin.Context) {
	name := c.Param("name")
	version := c.Param("version")

	for _, group := range storage.Groups {
		if group.Name == name && group.Version == version {
			c.JSON(http.StatusOK, group)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}

func GetGroupByID(c *gin.Context) {
	id := c.Param("id")

	group, exists := storage.Groups[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Group not found",
		})
		return
	}

	c.JSON(http.StatusOK, group)
}

func DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	_, exists := storage.Groups[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Group not found",
		})
		return
	}

	delete(storage.Groups, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Group deleted successfully",
	})
}

func DeleteGroupByVersion(c *gin.Context) {
	name := c.Param("name")
	version := c.Param("version")

	for id, group := range storage.Groups {
		if group.Name == name && group.Version == version {
			delete(storage.Groups, id)

			c.JSON(http.StatusOK, gin.H{
				"message": "Group deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}

func AddConfigurationToGroup(c *gin.Context) {
	groupID := c.Param("id")
	configID := c.Param("configId")

	group, groupExists := storage.Groups[groupID]
	if !groupExists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Group not found",
		})
		return
	}

	_, configExists := storage.Configurations[configID]
	if !configExists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Configuration not found",
		})
		return
	}

	for _, existingConfigID := range group.Configurations {
		if existingConfigID == configID {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Configuration already exists in group",
			})
			return
		}
	}

	group.Configurations = append(group.Configurations, configID)
	storage.Groups[groupID] = group

	c.JSON(http.StatusOK, group)
}

func RemoveConfigurationFromGroup(c *gin.Context) {
	groupID := c.Param("id")
	configID := c.Param("configId")

	group, groupExists := storage.Groups[groupID]
	if !groupExists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Group not found",
		})
		return
	}

	newConfigurations := []string{}
	found := false

	for _, existingConfigID := range group.Configurations {
		if existingConfigID == configID {
			found = true
			continue
		}
		newConfigurations = append(newConfigurations, existingConfigID)
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Configuration not found in group",
		})
		return
	}

	group.Configurations = newConfigurations
	storage.Groups[groupID] = group

	c.JSON(http.StatusOK, group)
}
