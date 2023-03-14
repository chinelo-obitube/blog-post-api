package models

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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