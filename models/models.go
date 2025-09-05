package models

import (
	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `json:"username" binding:"required" gorm:"unique"`
	Password   string `json:"password" binding:"required"`
	Blogs      []Blog `json:"blogs"`
	ProfilePic string `json:"profile_pic"`
}

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Picture     string `json:"picture"`
	UserID      uint   `json:"user_id"`
}

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	db.AutoMigrate(
		&User{},
		&Blog{},
	)

	DB = db
}
