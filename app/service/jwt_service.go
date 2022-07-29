package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatalln("secretKey not configured")
	}
	return secretKey
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
