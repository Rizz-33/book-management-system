package models

import (
	"github.com/Rizz-33/book-management-system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var books []Book
    db.Find(&books)
    return books
}

func GetBook(id int) Book {
	var book Book
    db.First(&book, id)
    return book
}

func CreateBook(book Book) {
    db.Create(&book)
}

func UpdateBook(book Book) {
    db.Save(&book)
}

func DeleteBook(id int) {
    db.Delete(&Book{}, id)
}