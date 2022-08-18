package middlewares

import (
	"boilerplate-api/api/services"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	jwtService services.JWTService
}

func NewJWTAuthMiddleware(
	jwtService services.JWTService,

) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		jwtService: jwtService,
	}
}

func (cc JWTAuthMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := cc.jwtService.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println("claims", claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Please enter a valid token string."})
		}
	}
}
