package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() *gorm.DB {
	database, err := gorm.Open("mysql", "Vectormike:password@/go_books_crud_dbms?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!", err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
