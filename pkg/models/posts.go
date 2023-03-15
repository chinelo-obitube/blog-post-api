package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/chinelo-obitube/goblog/pkg/config"
)

var db *gorm.DB

type BlogPost struct {
	ID int 	`gorm:"primary key" json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	CategoryID int `json:"categoryid"`
  Category string `json:"category"`
	User string `json:"user"`
}

type Category struct {
	ID int
	Name string
}

type User struct {
	ID int
	Name string
}


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BlogPost{}, &Category{}, &User{})
	fmt.Println("? Migration complete")
}


func (b *BlogPost) CreatePost() *BlogPost {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}