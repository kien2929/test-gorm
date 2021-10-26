package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@Admin12345@tcp(mysql:3306)/vap"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Migration() {
	User = 
	DB.Migrator().HasTable(&User{})
}
