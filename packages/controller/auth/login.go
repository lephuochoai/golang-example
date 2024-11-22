package controller_auth

import (
	"example/web-service-gin/packages/helper"
	"example/web-service-gin/packages/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login controller
func Login(c *gin.Context) {
	var input model.Authentication
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExist, err := model.FindOneUserBy(model.User{Email: input.Email})
	if (err != nil) || (userExist.ID == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errValidate := userExist.ValidatePassword(input.Password)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidate.Error()})
		return
	}

	accessToken, errGenJWT := helper.GenerateJWT(userExist)
	if errGenJWT != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errGenJWT.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
