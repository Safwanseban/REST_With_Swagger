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
// Create Books godoc
//	@Summary	creating a book
//	@Description
//	@Tags		book
//	@Accept		json
//	@Produce	json
//	@Success	200		{object} 	object
//	@Router		/create [post]
func CreateBooks(c *gin.Context) {
	var book models.Books
	books = append(books, book) //check the main function

	fmt.Println(books)
	c.JSON(200, gin.H{
		"msg": "created",
		"books":book,
	})
}
// Get Books godoc
//	@Summary	Getiing books
//	@Description
//	@Tags		book
//	@Accept		json
//	@Produce	json
//	@Success	200		{object} 	object
//	@Router		/get [get]
func Getbooks(c *gin.Context) {

	c.JSON(200, gin.H{
		"books": books,
	})
}
