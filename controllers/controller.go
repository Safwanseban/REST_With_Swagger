package controllers

import (
	"fmt"

	"github.com/Safwanseban/REST_With_Swagger/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "welcome home",
	})
}

var books []models.Books

func CreateBooks(c *gin.Context) {
	var book models.Books
	books = append(books, book) //check the main function

	fmt.Println(books)
	c.JSON(200, gin.H{
		"msg": "created",
	})
}
func Getbooks(c *gin.Context) {

	c.JSON(200, gin.H{
		"books": books,
	})
}
