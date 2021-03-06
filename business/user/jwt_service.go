package user

import (
	"fmt"
	"os"
	"time"

	"github.com/mashbens/cps/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string, ctx echo.Context) *jwt.Token
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "admin",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	config := config.GetConfig()
	jwtKey := config.App.JWTKey
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = jwtKey
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string, ctx echo.Context) *jwt.Token {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil
	}

	return t

}
