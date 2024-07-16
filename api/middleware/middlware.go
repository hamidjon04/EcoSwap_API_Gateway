package middleware

import (
	"api_gateway/api/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func JWTMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		auth := c.GetHeader("Authorization")

		if auth == ""{
			c.JSON(http.StatusBadRequest, fmt.Errorf("authorization header is required"))
			return
		}

		valid, err := token.ValidToken(auth)
		if !valid || err != nil{
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid token: %v", err))
			return 
		}

		claims, err := token.ExtraClaims(auth)
		if err != nil || !valid{
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid token claims: %v", err))
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}