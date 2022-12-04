package routes

import (
	"github.com/Safwanseban/REST_With_Swagger/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine) {
	c.GET("/", controllers.Home)
}
