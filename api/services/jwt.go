package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// type JWTService interface {
// 	GenerateToken(email string, isUser bool) string
// 	ValidateToken(token string) (*jwt.Token, error)
// }
type authCustomClaims struct {
	Email string `json:"email"`
	User  bool   `json:"user"`
	jwt.StandardClaims
}

type JWTService struct {
	secretKey string
	issure    string
}

// //auth-jwt
func NewJWTAuthService() JWTService {
	return JWTService{
		secretKey: getSecretKey(),
		issure:    "Prakash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (jwtService *JWTService) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    jwtService.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//String is encoded here
	t, err := token.SignedString([]byte(jwtService.secretKey))
	// fmt.Println("Genereated token", t)
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtService *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})
}
