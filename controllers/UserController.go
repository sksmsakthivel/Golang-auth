package controllers

import (
	"encoding/json"
	"fmt"
	"goauth/auth"
	"goauth/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/* signin the log in credentials*/
func SignIn(c *gin.Context) {

	userName, Password := c.PostForm("userName"), c.PostForm("password")

	if userName == "" || Password == "" {
		c.JSON(400, gin.H{"status": 0, "message": "userName and password is required"})
		return
	}
	user, err := models.FindUser(map[string]interface{}{"user_name": userName})
	if err != nil {
		c.JSON(400, gin.H{"status": 0, "message": "you are not registered with us"})
		return
	}
	fmt.Println("pass:", user, Password)
	if auth.CheckPasswordHash(Password, user.Password) {
		c.JSON(400, gin.H{"status": 0, "message": "your password is incorrect"})
		return
	}
	token, _ := auth.CreateToken(uint32(user.Id))

	c.JSON(200, gin.H{"status": 1, "message": "successfully login", "accesstoken": token})

}

/*signup the register the user*/
func SignUp(c *gin.Context) {
	var user models.User
	json.NewDecoder(c.Request.Body).Decode(&user)
	var validateErr = models.Validate(user)
	Password := c.PostForm("password")
	if Password != "" {
		validateErr = append(validateErr, "Password is required")
	}
	if len(validateErr) != 0 {
		c.JSON(400, gin.H{"status": 0, "message": validateErr})
		return
	}
	user.Password, _ = auth.HashPassword(Password)
	user.DOB, _ = time.Parse("02-01-2006", user.DateOfBirth)

	succesUser, err := models.CreateUser(&user)
	if err != nil {
		c.JSON(400, gin.H{"status": 0, "message": "Unable to regiter"})
		return
	}
	succesUser.DateOfBirth = user.DOB.Format("02-01-2006")
	c.JSON(200, gin.H{"status": 1, "message": "New Registration Successfull", "data": succesUser})
}

/*get user data by id*/
func GetUserById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindUser(map[string]interface{}{"id": Id})
	if err != nil {
		c.JSON(400, gin.H{"status": 0, "message": "you are not registered with us"})
		return
	}
	c.JSON(200, gin.H{"status": 1, "message": "get the user details", "data": user})
}
