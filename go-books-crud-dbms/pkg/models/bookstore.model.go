package models

import (
	"github.com/Vectormike/go-books-crud-dbms/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string  `gorm:""json:"name"`
	Author    string  `json:"author"`
	Price     float32 `json:"price"`
	Published bool    `json:"published"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&BookStore{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetBooks() []*Book {
	var books []*Book
	db.Find(&books)
	return books
}

func GetBook(id string) *Book {
	var book Book
	db.First(&book, id)
	return &book
}

func DeleteBook(id string) *Book {
	var book Book
	db.First(&book, id)
	db.Delete(&book)
	return &book
}
