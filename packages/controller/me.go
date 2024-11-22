package controller

import (
	"example/web-service-gin/packages/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	user := c.Keys["user"].(*model.User)
	newUserNoPass := model.RemovePassword(*user)
	c.JSON(http.StatusOK, gin.H{"user": newUserNoPass})
}
