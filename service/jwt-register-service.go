package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JWTRegisterService interface {
	GenerateToken(login string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type RegisterClaims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

type jwtRegisterService struct {
	secretKey string
	issuer    string
}

func NewJWTRegisterService() JWTRegisterService {
	return &jwtRegisterService{
		secretKey: getRegSecretKey(),
		issuer:    "makjac.pl",
	}
}

func getRegSecretKey() string {
	secret := os.Getenv("JWT_REGSEC")
	if secret == "" {
		secret = "bnm123"
	}
	return secret
}

func (jwtSrv *jwtRegisterService) GenerateToken(login string) (string, error) {
	//custom climes
	claims := &RegisterClaims{
		login,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 150).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		log.Printf("Error while generating token: %v", err)
		return "", err
	}
	return t, nil
}

func (jwtSrv *jwtRegisterService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}

func HashPasswd(passwd string) (string, error) {
	bytePW := []byte(passwd)
	hashPW, err := bcrypt.GenerateFromPassword(bytePW, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hash: %v", err)
		return "", err
	}
	return string(hashPW), nil
}
