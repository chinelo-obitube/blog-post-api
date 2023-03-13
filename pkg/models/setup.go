package models

import (
  "gorm.io/gorm"
  _ "gorm.io/driver/sqlite"
	"github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func connectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})


        if err != nil {
                panic("Failed to connect to database!")
        }

        db = database
}

func GetDB() *gorm.DB{
	return db
}