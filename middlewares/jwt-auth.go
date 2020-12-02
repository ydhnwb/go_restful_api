package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := entities.BuildErrorResponse("Failed to process your request", "Cannot get token from Authorization herader", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Issuer]: ", claims["issuer"])
			log.Println("Claims[ExpiredAt]: ", claims["exp"])
		} else {
			log.Println(err)
			response := entities.BuildErrorResponse("Failed to validate your token ", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}
}
