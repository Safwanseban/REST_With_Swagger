package main

import (
	"github.com/Safwanseban/REST_With_Swagger/configs"
	"github.com/Safwanseban/REST_With_Swagger/models"
	"github.com/Safwanseban/REST_With_Swagger/routes"
	"github.com/Safwanseban/REST_With_Swagger/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

var env *configs.EnvConfig
var R = gin.Default()

func init() {
	env = configs.Envload()
	// fmt.Println(env.Port)
}
var books []models.Books
func main() {
	docs.SwaggerInfo.Title = "E-Commerce API"
	docs.SwaggerInfo.Description = "An e-commerce API which is purely written in GO."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	books = append(books, models.Books{Book_ID: 1, Book_name: "Wings of fire", Author_name: "APJ Abdul-Kalam"})
	books = append(books, models.Books{Book_ID: 2, Book_name: "The Story of My Experiments with Truth",
		Author_name: "Mahatma Gandhi"})
	routes.Routes(R)
	R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	R.Run(env.Port)
}
