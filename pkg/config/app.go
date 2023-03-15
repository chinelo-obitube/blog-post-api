package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)


func Connect() {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println("Failed to connect to the database")
 }

defer db.Close()


}

func GetDB() *gorm.DB{
	return db
}