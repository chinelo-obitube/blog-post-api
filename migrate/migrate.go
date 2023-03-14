package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/chinelo-obitube/goblog/models"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
}
defer db.Close()

  db.AutoMigrate(&models.BlogPost{}, &models.Category{}, &models.User{})
	fmt.Println("? Migration complete")

}