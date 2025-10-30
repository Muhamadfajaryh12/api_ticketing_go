package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorization() gin.HandlerFunc{
	return func(c *gin.Context){
		header := c.GetHeader("Authorization")

		if header == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": true,"message":"Token tidak di temukan"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(header,"Bearer ")
		secret:= os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func (t *jwt.Token)(interface{},error){
			return []byte(secret),nil
		})

		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"error": true,"message":"Token tidak valid"})
			c.Abort()
			return
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("name", claims["name"])
			c.Set("role",claims["role"])
		}

		c.Next()
	}
}