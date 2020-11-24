package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/services"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT() gin.HandlerFunc{
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenGiven := authHeader[len(BEARER_SCHEMA):]
		token, err := services.NewJWTService().ValidateToken(tokenGiven)
		if token.Valid {
			claims := 
		}

	}

}