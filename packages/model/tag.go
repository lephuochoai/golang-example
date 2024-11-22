package model

import (
	"example/web-service-gin/packages/database"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
	Task []Task `json:"tasks" gorm:"many2many:task_tags;"`
}

func (tag *Tag) Save() (*Tag, error) {
	err := database.Database.Create(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func FindOneTagBy(query interface{}) (*Tag, error) {
	var tag *Tag
	err := database.Database.Where(query).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}
