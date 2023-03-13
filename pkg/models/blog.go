package models

import (
		"gorm.io/gorm"
		_"gorm.io/gorm/sqlite"
)

type post struct {
	ID uint `json:"id" gorm:"primary key"`
	Title string `json:"title"`	
  Body string `json:"body"`
	Category string `json:"category"`
}