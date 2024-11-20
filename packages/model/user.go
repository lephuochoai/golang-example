package model

import (
	"example/web-service-gin/packages/database"
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func FindOneBy(query interface{}) (User, error) {
	var user User
	err := database.Database.Where(query).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user *User) ValidatePassword(password string) error {
	fmt.Println([]byte(password))
	fmt.Println([]byte(user.Password))
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
