package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

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

func main() {
	db, _ = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
}
defer db.Close()

  db.AutoMigrate(&BlogPost{}, &Category{}, &User{})

	return db
}

router := gin.Default()
router.GET("/post", GetAllPosts)
router.GET("/post/:id", GetPostById)
router.PUT("/post/:id", UpdatePost)
router.POST("/post/", CreatePost)
router.DELETE("/post/:id", DeletePost)
router.Run(":8001")