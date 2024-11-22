package controller

import (
	"example/web-service-gin/packages/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Projects(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	limitStr := queryParams.Get("limit")
	pageStr := queryParams.Get("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	projects, err := model.Project{}.FindPaginate(limit, page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func CreateProject(c *gin.Context) {
	var input model.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := input.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": project})
}
