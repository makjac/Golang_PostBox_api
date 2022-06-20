package middlewares

import (
	"log"
	"net/http"

	"github.com/PBB-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHader := c.GetHeader("Authorization")

		if authHader == "" {
			log.Println("no token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "you are not loged in",
			})
			c.Abort()
			return
		}

		tokenString := authHader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("uuid", claims["uuid"])
			c.Next()
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
