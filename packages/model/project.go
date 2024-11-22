package model

import (
	"example/web-service-gin/packages/database"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks" gorm:"foreignKey:ProjectId;"`
}

func (project Project) FindPaginate(limit int, page int) ([]Project, error) {
	var projects []Project
	projectsResult := database.Database.Limit(limit).Offset((page - 1) * limit).Find(&projects)

	if projectsResult.Error != nil {
		return nil, projectsResult.Error
	}
	return projects, nil
}

func (project *Project) Save() (*Project, error) {
	err := database.Database.Create(&project).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}

func FindOneProjectBy(query interface{}) (*Project, error) {
	var project *Project
	err := database.Database.Where(query).First(&project).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}
