package models

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open("postgresql://postgres:1234@localhost:5432/GO_CRUD"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&User{})
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserReal struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var Users []UserReal
var NextID = 1
