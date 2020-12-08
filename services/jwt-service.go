package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

//Extending the jwt.standarClaim, use this for learning purpose
type jwtCustomClaim struct {
	UserID string `json:"user_id"`
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
	secretKey := os.Getenv("JWT_SECRETKEY")
	if secretKey == "" {
		secretKey = "exampleOfMySecretKey"
	}
	return secretKey
}

func (jwtService *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
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
