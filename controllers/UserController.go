package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	fmt.Println("signuin")
	json.NewDecoder(c.Request.Body)
}
func SignUp(c *gin.Context) {
	fmt.Println("signuup")
	json.NewDecoder(c.Request.Body)
}
