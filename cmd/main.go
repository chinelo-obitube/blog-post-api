package main

import (
	"fmt"
	"logs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

func main(){
	router := gin.Default(
	router.GET("/posts", getPosts)
	router.GET("/posts/{id}", getPostsByID)
	router.GET("/posts/{category}",getPostsByCategory)
	router.POST("/posts/{category}", CreatePost)
	router.PUT("/posts/{category}", UpdatePost)
	router.DELETE("/posts/{id}", DeletePost)
	)
}