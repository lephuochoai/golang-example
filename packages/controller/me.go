package controller

import (
	"example/web-service-gin/packages/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserNoPass struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Me controller
func Me(c *gin.Context) {
	userId := c.Keys["userId"].(model.User).ID
	user, err := model.FindOneBy(model.User{ID: userId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newUserNoPass UserNoPass
	newUserNoPass.ID = user.ID
	newUserNoPass.Name = user.Name
	newUserNoPass.Email = user.Email

	c.JSON(http.StatusOK, gin.H{"user": newUserNoPass})
}
