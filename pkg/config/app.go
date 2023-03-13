package config

import(
	"gorm.io/gorm"
	_"gorm.io/gorm/sqlite"
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