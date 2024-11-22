package model

import (
	"example/web-service-gin/packages/database"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	DueTask     string `json:"due_task"`
	Priority    string `json:"priority"`
	ProjectId   uint   `json:"project_id"`
	Project     Project
	Tags        []Tag `json:"tags" gorm:"many2many:task_tags;"`
}

func (task *Task) Save() (*Task, error) {
	err := database.Database.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (task Task) FindPaginate(limit int, page int) ([]Task, error) {
	var tasks []Task
	tasksResult := database.Database.Preload("Tags").Preload("Project").Limit(limit).Offset((page - 1) * limit).Find(&tasks)

	if tasksResult.Error != nil {
		return nil, tasksResult.Error
	}
	return tasks, nil
}
