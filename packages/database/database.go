package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)

	fmt.Println(dns)

	Db, err := gorm.Open(
		mysql.Open(dns), &gorm.Config{})

	Database = Db

	if err != nil {
		panic(err)
	}
}
