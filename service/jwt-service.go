package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(username string, uuid string) string
	ValidateToken(tokenStrring string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "makjac.pl",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "123bnm"
	}
	return secret
}

func (jwtSrv *jwtService) GenerateToken(username string, uuid string) string {

	//custom climes
	claims := &jwtCustomClaims{
		username,
		uuid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	//create token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//generate encoded token by secret ket
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		log.Printf("Error while generating token: %v", err)
	}
	return t
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
