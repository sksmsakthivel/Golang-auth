package auth

import (
	"fmt"
	"goauth/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*create JWT token*/
func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))

}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware
		token := c.GetHeader("Authorization")
		fmt.Println("token", token)
		if token == "" {
			c.Abort()
			c.JSON(400, gin.H{"status": 0, "message": "Token is requireds"})
			return
		}

		if len(strings.Split(token, " ")) == 2 {
			if strings.Split(token, " ")[0] != "Bearer" {
				fmt.Println("bearer")
				c.Abort()
				c.JSON(400, gin.H{"status": 0, "message": "Token is invalid"})
				return
			}
			tokenString := strings.Split(token, " ")[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.Abort()
					c.JSON(400, gin.H{"status": 0, "message": "Token is invalid"})
					// return

					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
			
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				c.Abort()
				c.JSON(400, gin.H{"status": 0, "message": "Token is invalid"})
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				uid, _ := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)

				_, err = models.FindUser(map[string]interface{}{"id": uid})
				if err != nil {
					c.Abort()
					c.JSON(401, gin.H{"status": 0, "message": "You acces denied"})
					return
				}
				c.Next()
			}
		} else {
			c.Abort()
			c.JSON(400, gin.H{"status": 0, "message": "Token is invalid"})
			return
		}
	}
}
