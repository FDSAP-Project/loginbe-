package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitialMigration() {
	var DNS = "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("not connected")
	} else {
		fmt.Println("connected")
	}
	Db.AutoMigrate(&User{})
}

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
