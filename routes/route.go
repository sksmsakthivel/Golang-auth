package routes

import (
	"goauth/auth"
	"goauth/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/signin", controllers.SignIn)
	r.POST("/signup", controllers.SignUp)
	r.Use(auth.Authorize()).GET("/get-user/:id", controllers.GetUserById)
	return r
}
