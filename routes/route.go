package routes

import (
	"goauth/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() {
	r := gin.Default()
	r.POST("/signin", controllers.SignIn)
	r.POST("/signup", controllers.SignUp)
	// return r
}
