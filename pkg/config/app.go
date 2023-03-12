package config

import(
	"github.co/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}