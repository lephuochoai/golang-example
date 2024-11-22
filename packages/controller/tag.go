package controller

import (
	"example/web-service-gin/packages/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagExist, err := model.FindOneTagBy(&model.Tag{Name: tag.Name})
	if err == nil || tagExist != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag already exists"})
		return
	}

	_, err = tag.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tag": tag})
}
