package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/chinelo-obitube/goblog/models"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	DB *gorm.DB
}

var NewPost models.Book

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func GetPost(context *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		context.JSON(200, people)
	}
}
