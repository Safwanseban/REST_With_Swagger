package main

import (
	"github.com/Safwanseban/REST_With_Swagger/configs"
	"github.com/Safwanseban/REST_With_Swagger/routes"
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

func main() {
	routes.Routes(R)
	R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	R.Run(env.Port)
}
