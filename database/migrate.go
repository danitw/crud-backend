package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"     gorm:"not null"`
	Email    string `json:"email"    gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"   gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	UserId  int    `json:"user_id" gorm:"not null"`
	User    User   `               gorm:"not null"`
}

func Migrate(db *gorm.DB) {
  db.Debug().AutoMigrate(&User{}, &Post{})
}
