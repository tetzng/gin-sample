package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)

	router.Run("localhost:8080")
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []post{
	{ID: "1", Title: "Hello Gin", Body: "I started to learn golang."},
	{ID: "2", Title: "Second Post", Body: "Second body."},
	{ID: "3", Title: "TITLE", Body: "MY BODY"},
}
