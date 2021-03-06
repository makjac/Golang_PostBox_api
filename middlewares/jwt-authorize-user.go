package middlewares

import (
	"log"
	"net/http"

	"github.com/PBB-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeUserJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var UUID = c.Param("uuid")

		const BEARER_SCHEMA = "Bearer "
		authHader := c.GetHeader("Authorization")
		tokenString := authHader[len(BEARER_SCHEMA):]

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "you are not loged in",
			})
			c.Abort()
			return
		}

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
		} else {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Wrong token",
			})
			c.Abort()
			return
		}
	}
}

func CheckUser(username string) bool {

	return false
}
