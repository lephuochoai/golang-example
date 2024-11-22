package controller

import (
	"example/web-service-gin/packages/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectExist, err := model.FindOneProjectBy(&model.Project{Model: gorm.Model{ID: task.ProjectId}})

	if err != nil || projectExist == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found"})
		return
	}

	if _, err := task.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Project = *projectExist

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func FindPaginateTask(c *gin.Context) {
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

	tasks, err := model.Task{}.FindPaginate(limit, page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
