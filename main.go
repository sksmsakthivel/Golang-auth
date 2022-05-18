package main

import (
	// "goauth/routes"
	"fmt"
	"goauth/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// routes.SetUpRoutes()
	r := gin.Default()
	r.GET("/signin", controllers.SignIn)
	r.GET("/signup", controllers.SignUp)
	r.Run()
	fmt.Println("hello")
}
