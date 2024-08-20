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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	  c.Writer.Header().Set("Content-Type", "application/json")
	  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	  c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
  
	  if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	  } else {
		c.Next()
	  }
	}
  }
  