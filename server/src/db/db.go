package db

import (
	model "server/graph/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:password@tcp(db)/tm_app_server_db?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	DB.AutoMigrate(&model.Todo{}, &model.User{})
}

func CloseDB() {
	database, err := DB.DB()
	if err != nil {
		panic(err)
	}
	database.Close()
}
