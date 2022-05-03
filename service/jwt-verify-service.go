package service

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTVerifyService interface {
	GenerateToken(email string) string
	ValidateVerifyToken(jwtToken string) (*jwt.Token, error)
}

type jwtVerifyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type jwtVerifyService struct {
	secretKey string
	issuer    string
}

func (jwtSrv *jwtVerifyService) GenerateToken(email string) string {

	//custom climes
	claims := &jwtCustomClaims{
		email,
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

func (jwtSrv *jwtService) ValidateVerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
