package model

import (
	"example/web-service-gin/packages/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive" gorm:"default:true"`
}

type UserNoPass struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	return nil
}

func FindOneUserBy(query interface{}) (*User, error) {
	var user *User
	err := database.Database.Where(query).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func RemovePassword(user User) UserNoPass {
	return UserNoPass{
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
}
