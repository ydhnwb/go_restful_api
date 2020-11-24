package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

//Extending the jwt.standarClaim, use this for learning purpose
type jwtCustomClaim struct {
	Name  string `json:"name"`
	Admin bool   `json:"is_admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("secretKey")
	if secretKey == "" {
		secretKey = "exampleOfMySecretKey"
	}
	return secretKey
}

func (jwtService *jwtService) GenerateToken(email string, admin bool) string {
	claims := &jwtCustomClaim{
		email,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtService.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtService *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(tokenForFunc *jwt.Token) (interface{}, error) {
		if _, ok := tokenForFunc.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", tokenForFunc.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})
}
