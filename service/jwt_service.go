package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTKey struct {
	SecretKey []byte
}

type JWTService interface {
	GenerateToken(name string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "dipto",
	}
}

type jwtClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		secret = "diptochuck"
	}
	return secret
}

func (jwtsvc *jwtService) GenerateToken(username string) string {
	claims := &jwtClaim{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtsvc.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSign, err := token.SignedString([]byte(jwtsvc.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenSign
}

func (jwtsvc *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("incorrect SignIn Method: %v", t.Header["alg"])
		}
		return []byte(jwtsvc.secretKey), nil
	})
}
