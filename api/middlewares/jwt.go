package middlewares

import (
	"boilerplate-api/api/services"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	service services.JWTService
}

func NewJWTAuthMiddleware(
	service services.JWTService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: service,
	}
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := services.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println("claims", claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Please enter a valid token string."})
		}
	}
}
