package controller_auth

import (
	"example/web-service-gin/packages/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register controller
func Register(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExist, err := model.FindOneUserBy(&model.User{Email: input.Email})

	if err == nil || userExist != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	_, err = input.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
